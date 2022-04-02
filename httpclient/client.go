package httpclient

import (
	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	"io"
	"net/http"
	"time"
)

var _ Client = (*client)(nil)

type Client interface {
	HTTPClient() *httpclient.Client
	GET(url string, headers http.Header) (*http.Response, error)
	POST(url string, body io.Reader, headers http.Header) (*http.Response, error)
	PUT(url string, body io.Reader, headers http.Header) (*http.Response, error)
	PATCH(url string, body io.Reader, headers http.Header) (*http.Response, error)
	DELETE(url string, headers http.Header) (*http.Response, error)
}

type Options httpclient.Option

// WithHTTPTimeout 设置超时时间
func WithHTTPTimeout(timeout time.Duration) Options {
	return Options(httpclient.WithHTTPTimeout(timeout))
}

// WithRetryCount 设置重试次数
func WithRetryCount(retryCount int) Options {
	return Options(httpclient.WithRetryCount(retryCount))
}

// WithRetrier 设置重试策略
func WithRetrier(retrier heimdall.Retriable) Options {
	return Options(httpclient.WithRetrier(retrier))
}

// WithHTTPClient 设置自定义HTTP客户端
func WithHTTPClient(client heimdall.Doer) Options {
	return Options(httpclient.WithHTTPClient(client))
}

type client struct {
	client *httpclient.Client
}

func NewClient(opts ...Options) Client {
	var c = new(client)
	var o []httpclient.Option
	for _, v := range opts {
		o = append(o, httpclient.Option(v))
	}
	c.client = httpclient.NewClient(o...)
	return c
}

func (c *client) HTTPClient() *httpclient.Client {
	return c.client
}

func (c *client) GET(url string, headers http.Header) (*http.Response, error) {
	return c.client.Get(url, headers)
}

func (c *client) POST(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	return c.client.Post(url, body, headers)
}

func (c *client) PUT(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	return c.client.Put(url, body, headers)
}

func (c *client) PATCH(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	return c.client.Patch(url, body, headers)
}

func (c *client) DELETE(url string, headers http.Header) (*http.Response, error) {
	return c.client.Delete(url, headers)
}

