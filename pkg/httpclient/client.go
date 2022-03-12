package httpclient

import (
	"github.com/gojek/heimdall/v7/httpclient"
	"io"
	"net/http"
)

var _ Client = (*client)(nil)

type Client interface {
	HttpClient() *httpclient.Client
	GET(url string, headers http.Header) (*http.Response, error)
	POST(url string, body io.Reader, headers http.Header) (*http.Response, error)
	PUT(url string, body io.Reader, headers http.Header) (*http.Response, error)
	PATCH(url string, body io.Reader, headers http.Header) (*http.Response, error)
	DELETE(url string, headers http.Header) (*http.Response, error)
}

type client struct {
	*httpclient.Client
}

func NewClient(opts ...httpclient.Option) Client {
	var client = new(client)
	client.Client = httpclient.NewClient(opts...)
	return client
}

func (c *client) HttpClient() *httpclient.Client {
	return c.Client
}

func (c *client) GET(url string, headers http.Header) (*http.Response, error) {
	return c.Client.Get(url, headers)
}

func (c *client) POST(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	return c.Client.Post(url, body, headers)
}

func (c *client) PUT(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	return c.Client.Put(url, body, headers)
}

func (c *client) PATCH(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	return c.Client.Patch(url, body, headers)
}

func (c *client) DELETE(url string, headers http.Header) (*http.Response, error) {
	return c.Client.Delete(url, headers)
}

