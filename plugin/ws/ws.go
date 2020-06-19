package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/gopub/errors"
	"github.com/gopub/log"
	"github.com/gopub/wine"
)

var logger = wine.Logger()

func SetLogger(l *log.Logger) {
	logger = l
}

type GetAuthUserID interface {
	GetAuthUserID() int64
}

const (
	methodGetDate = "ws.getDate"
)

func MarshalData(v interface{}) (*Data, error) {
	if d, ok := v.(*Data); ok {
		return d, nil
	}
	d := new(Data)
	if b, ok := v.([]byte); ok {
		d.V = &Data_Raw{
			Raw: b,
		}
	} else if m, ok := v.(proto.Message); ok {
		b, err := proto.Marshal(m)
		if err != nil {
			return nil, err
		}
		d.V = &Data_Protobuf{
			Protobuf: b,
		}
	} else {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		d.V = &Data_Json{
			Json: b,
		}
	}
	return d, nil
}

func (m *Data) Unmarshal(v interface{}) error {
	switch dv := m.V.(type) {
	case *Data_Json:
		return json.Unmarshal(dv.Json, v)
	case *Data_Protobuf:
		if m, ok := v.(proto.Message); ok {
			return proto.Unmarshal(dv.Protobuf, m)
		}
		return fmt.Errorf("v is %T not of proto.Message type", v)
	case *Data_Raw:
		val := reflect.ValueOf(v)
		if val.Kind() != reflect.Ptr || val.Elem().Type() != reflect.SliceOf(reflect.TypeOf(byte(0))) {
			return fmt.Errorf("cannot unmarshal []byte into %T", v)
		}
		val.Elem().SetBytes(dv.Raw)
		return nil
	default:
		return fmt.Errorf("cannot unmarshal data type: %T", v)
	}
}

func NewCall(id int32, name string, params interface{}) (*Call, error) {
	data, err := MarshalData(params)
	if err != nil {
		return nil, err
	}
	return &Call{
		Id:   id,
		Name: name,
		Data: data,
	}, nil
}

func NewReply(id int32, resultOrErr interface{}) *Reply {
	r := new(Reply)
	r.Id = id
	if err, ok := resultOrErr.(error); ok {
		code := errors.GetCode(err)
		if code <= 0 {
			code = http.StatusInternalServerError
		}
		r.Result = &Reply_Error{
			Error: &Error{
				Code:    int32(code),
				Message: err.Error(),
			},
		}
	} else {
		data, err := MarshalData(resultOrErr)
		if err != nil {
			r.Result = &Reply_Error{
				Error: &Error{
					Code:    int32(http.StatusInternalServerError),
					Message: err.Error(),
				},
			}
		} else {
			r.Result = &Reply_Data{
				Data: data,
			}
		}
	}
	return r
}

func NewDataPacket(v interface{}) (*Packet, error) {
	data, err := MarshalData(v)
	if err != nil {
		return nil, err
	}
	p := new(Packet)
	p.V = &Packet_Data{
		Data: data,
	}
	return p, nil
}

func NewPushPacket(typ int32, data interface{}) (*Packet, error) {
	d, err := MarshalData(data)
	if err != nil {
		return nil, err
	}
	p := new(Packet)
	p.V = &Packet_Push{
		Push: &Push{
			Type: typ,
			Data: d,
		},
	}
	return p, nil
}

type PacketReadWriter interface {
	Read() (*Packet, error)
	Write(p *Packet) error
}

func (m *Packet) UnmarshalData(v interface{}) error {
	switch v := m.V.(type) {
	case *Packet_Data:
		if v.Data != nil {
			return v.Data.Unmarshal(v)
		}
	case *Packet_Reply:
		switch res := v.Reply.Result.(type) {
		case *Reply_Data:
			return res.Data.Unmarshal(v)
		}
	}
	return nil
}

type contextKey int

// Context keys
const (
	ckNextHandler contextKey = iota + 1
	ckAuthFlag
	ckPusher
)

type Pusher interface {
	Push(ctx context.Context, userID int64, typ int32, data interface{}) error
}

func GetPusher(ctx context.Context) Pusher {
	p, _ := ctx.Value(ckPusher).(Pusher)
	return p
}

func withPusher(ctx context.Context, p Pusher) context.Context {
	return context.WithValue(ctx, ckPusher, p)
}
