package ws

import (
	"container/list"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gopub/errors"

	"github.com/gorilla/websocket"
)

type ClientState int

const (
	Disconnected ClientState = iota
	Connecting
	Connected
	Closed
)

type Client struct {
	dialTimeout      time.Duration
	pingInterval     time.Duration
	maxReconnBackoff time.Duration
	reconnBackoff    time.Duration

	addr string

	newCallC chan struct{}
	mu       sync.RWMutex // guard calls, replyM
	calls    *list.List
	replyM   map[int32]chan<- *Reply
	header   map[string]string

	conn   *Conn
	state  ClientState
	stateC chan ClientState

	callID int32

	Handshake func(rw PacketReadWriter) error
	pushDataC chan []byte

	CallLogger func(call *Call, reply *Reply, callAt time.Time)
}

func NewClient(addr string) *Client {
	c := &Client{
		dialTimeout:      10 * time.Second,
		pingInterval:     10 * time.Second,
		maxReconnBackoff: 2 * time.Second,
		addr:             addr,
		calls:            list.New(),
		newCallC:         make(chan struct{}, 1),
		replyM:           make(map[int32]chan<- *Reply),
		state:            Disconnected,
		stateC:           make(chan ClientState, 1),
		pushDataC:        make(chan []byte, 1),
		callID:           1,
		header:           map[string]string{},
	}
	c.CallLogger = c.logCall
	go c.start()
	return c
}

func (c *Client) nextCallID() int32 {
	atomic.AddInt32(&c.callID, 2)
	return c.callID
}

func (c *Client) start() {
	c.reconnBackoff = 100 * time.Millisecond
	for c.state != Closed {
		c.run()
		if c.reconnBackoff > 0 {
			time.Sleep(c.reconnBackoff)
		}
		c.reconnBackoff += 100 * time.Millisecond
		if c.reconnBackoff > c.maxReconnBackoff {
			c.reconnBackoff = c.maxReconnBackoff
		}
	}
}

func (c *Client) run() {
	c.setState(Connecting)
	ctx, cancel := context.WithTimeout(context.Background(), c.dialTimeout)
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, c.addr, nil)
	if err != nil {
		cancel()
		logger.Errorf("Cannot connect %s: %v", c.addr, err)
		c.setState(Disconnected)
		return
	}
	cancel()
	c.conn = NewConn(conn)
	if c.Handshake != nil {
		if err = c.Handshake(c.conn); err != nil {
			logger.Errorf("Cannot handshake: %v", err)
			c.setState(Disconnected)
			return
		}
	}
	if err = c.writeHeader(); err != nil {
		c.setState(Disconnected)
		return
	}
	c.setState(Connected)
	done := make(chan struct{}, 1)
	go c.read(done)
	c.write(done)
	c.setState(Disconnected)
}

func (c *Client) read(done chan<- struct{}) {
	defer logger.Debug("Exited read loop")
	for {
		p, err := c.conn.Read()
		if err != nil {
			logger.Errorf("Cannot read: %v", err)
			done <- struct{}{}
			return
		}
		if v, ok := p.V.(*Packet_Push); ok {
			select {
			case c.pushDataC <- v.Push:
				break
			default:
				break
			}
		}
		switch v := p.V.(type) {
		case *Packet_Push:
			select {
			case c.pushDataC <- v.Push:
				break
			default:
				break
			}
		case *Packet_Reply:
			c.mu.RLock()
			if ch, ok := c.replyM[v.Reply.Id]; ok {
				ch <- v.Reply
				delete(c.replyM, v.Reply.Id)
			}
			c.mu.RUnlock()
		}
	}
}

func (c *Client) write(done <-chan struct{}) {
	defer logger.Debug("Exited write loop")
	t := time.NewTicker(c.pingInterval)
	m := NewNetworkMonitor()
	defer m.Stop()
	defer t.Stop()
	for {
		select {
		case <-t.C:
			if err := c.conn.Push(nil); err != nil {
				logger.Errorf("Cannot ping: %v", err)
				c.reconnBackoff = 0
				return
			}
			logger.Debugf("Ping")
		case <-m.C:
			c.reconnBackoff = 0
			return
		case <-done:
			c.reconnBackoff = 0
			return
		case <-c.newCallC:
			c.mu.Lock()
		CallLoop:
			for it := c.calls.Front(); it != nil; {
				ca := it.Value.(*Call)
				next := it.Next()
				c.calls.Remove(it)
				it = next
				if err := c.conn.Call(ca); err != nil {
					logger.Errorf("Cannot call %s: %v", ca.Name, err)
					if rc, ok := c.replyM[ca.Id]; ok {
						select {
						case rc <- NewErrorReply(ca.Id, err):
							break
						default:
							break
						}
						delete(c.replyM, ca.Id)
					}
					break CallLoop
				}
			}
			c.mu.Unlock()
		}
	}
}

func (c *Client) writeHeader() error {
	if len(c.header) == 0 {
		return nil
	}
	h := &Header{
		Entries: c.header,
	}
	p := new(Packet)
	p.V = &Packet_Header{
		Header: h,
	}
	return c.conn.Write(p)
}

func (c *Client) Call(ctx context.Context, name string, params interface{}, result interface{}) error {
	if c.state == Closed {
		return errors.New("client is closed")
	}
	data, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("cannot marshal params: %w", err)
	}
	ca := NewCall(c.nextCallID(), name, data)
	replyC := make(chan *Reply, 1)
	c.mu.Lock()
	c.calls.PushBack(ca)
	c.replyM[ca.Id] = replyC
	c.mu.Unlock()
	select {
	case c.newCallC <- struct{}{}:
		break
	default:
		break
	}

	startAt := time.Now()
	select {
	case <-ctx.Done():
		if c.CallLogger != nil {
			c.CallLogger(ca, NewErrorReply(ca.Id, ctx.Err()), startAt)
		}
		return ctx.Err()
	case reply := <-replyC:
		if c.CallLogger != nil {
			c.CallLogger(ca, reply, startAt)
		}
		switch v := reply.Result.(type) {
		case *Reply_Data:
			if result == nil {
				break
			}
			if len(v.Data) == 0 {
				return errors.New("no data")
			}
			if err := json.Unmarshal(v.Data, result); err != nil {
				return fmt.Errorf("cannot unmarshal result: %w", err)
			}
		case *Reply_Error:
			return errors.Format(int(v.Error.Code), v.Error.Message)
		}
		return nil
	}
}

func (c *Client) SetHeader(h map[string]string) {
	c.mu.Lock()
	for k, v := range h {
		c.header[k] = v
	}
	c.mu.Unlock()
	if c.state == Connected {
		go c.writeHeader()
	}
}

func (c *Client) Close() {
	c.setState(Closed)
	close(c.pushDataC)
	close(c.stateC)
}

func (c *Client) setState(s ClientState) {
	if c.state == s || c.state == Closed {
		return
	}
	c.state = s
	switch s {
	case Disconnected, Closed:
		if c.conn != nil {
			c.conn.Close()
			c.conn = nil
		}
	}
	select {
	case c.stateC <- s:
		break
	default:
		break
	}
}

func (c *Client) State() ClientState {
	return c.state
}

func (c *Client) StateC() <-chan ClientState {
	return c.stateC
}

func (c *Client) SetDialTimeout(t time.Duration) {
	if t < time.Second {
		t = time.Second
	}
	c.dialTimeout = t
}

func (c *Client) SetPingInterval(t time.Duration) {
	if t < time.Second {
		t = time.Second
	}
	c.pingInterval = t
}

func (c *Client) SetMaxReconnBackoff(t time.Duration) {
	if t <= 0 {
		t = 0
	}
	c.maxReconnBackoff = t
}

func (c *Client) PushDataC() <-chan []byte {
	return c.pushDataC
}

func (c *Client) GetServerTime(ctx context.Context) (time.Time, error) {
	var res struct {
		Timestamp int64 `json:"timestamp"`
	}
	err := c.Call(ctx, methodGetDate, nil, &res)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(res.Timestamp, 0), nil
}

func (c *Client) logCall(call *Call, reply *Reply, callAt time.Time) {
	cost := time.Since(callAt)
	switch v := reply.Result.(type) {
	case *Reply_Data:
		logger.Infof("%d %s | %v", call.Id, call.Name, cost)
	case *Reply_Error:
		logger.Errorf("%d %s | %s | %d:%s | %v", call.Id, call.Name, call.Data, v.Error.Code, v.Error.Message, cost)
	}
}
