package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gopub/gox"
	"github.com/gopub/wine/mime"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type HeaderBuilder interface {
	Build(ctx context.Context, header http.Header) http.Header
}

type Client struct {
	client        *http.Client
	header        http.Header
	HeaderBuilder HeaderBuilder
}

func NewClient(client *http.Client) *Client {
	return &Client{
		client: client,
		header: make(http.Header),
	}
}

func (c *Client) HTTPClient() *http.Client {
	return c.client
}

func (c *Client) Header() http.Header {
	return c.header
}

func (c *Client) InjectHeader(req *http.Request) {
	for k, vs := range c.header {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}
	if c.HeaderBuilder != nil {
		req.Header = c.HeaderBuilder.Build(req.Context(), req.Header)
	}
}

func (c *Client) Get(ctx context.Context, endpoint string, query url.Values, result interface{}) error {
	if query == nil {
		return c.call(ctx, http.MethodGet, endpoint, nil, result)
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return errors.Wrap(err, "parse url failed")
	}
	q := u.Query()
	for k, vs := range query {
		for _, v := range vs {
			q.Add(k, v)
		}
	}
	u.RawQuery = q.Encode()
	return c.call(ctx, http.MethodGet, u.String(), nil, result)
}

func (c *Client) Post(ctx context.Context, endpoint string, params interface{}, result interface{}) error {
	return c.call(ctx, http.MethodPost, endpoint, params, result)
}

func (c *Client) Put(ctx context.Context, endpoint string, params interface{}, result interface{}) error {
	return c.call(ctx, http.MethodPut, endpoint, params, result)
}

func (c *Client) Patch(ctx context.Context, endpoint string, params interface{}, result interface{}) error {
	return c.call(ctx, http.MethodPatch, endpoint, params, result)
}

func (c *Client) Delete(ctx context.Context, endpoint string, query url.Values, result interface{}) error {
	if query == nil {
		return c.call(ctx, http.MethodDelete, endpoint, nil, result)
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return errors.Wrap(err, "parse url failed")
	}
	q := u.Query()
	for k, vs := range query {
		for _, v := range vs {
			q.Add(k, v)
		}
	}
	u.RawQuery = q.Encode()
	return c.call(ctx, http.MethodDelete, endpoint, u.String(), result)
}

func (c *Client) call(ctx context.Context, method string, endpoint string, params interface{}, result interface{}) error {
	var body io.Reader
	if params != nil {
		data, err := json.Marshal(params)
		if err != nil {
			return errors.Wrap(err, "marshal failed")
		}
		body = bytes.NewBuffer(data)
	}
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return errors.Wrap(err, "create request failed")
	}
	if params != nil {
		req.Header.Set(mime.ContentType, mime.JSON)
	}
	req = req.WithContext(ctx)
	return c.Do(req, result)
}

func (c *Client) Do(req *http.Request, result interface{}) error {
	c.InjectHeader(req)
	resp, err := c.client.Do(req)
	if err != nil {
		return gox.NewError(StatusTransportFailed, err.Error())
	}
	return ParseResult(resp, result)
}
