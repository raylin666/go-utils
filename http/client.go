package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	ut "github.com/raylin666/go-utils"
)

var _ Client = (*client)(nil)

var (
	ErrUrlEmpty          = fmt.Errorf("URL cannot be empty")
	ErrUrlInvalid        = fmt.Errorf("URL format invalid")
	ErrTimeoutInvalid    = fmt.Errorf("timeout must be greater than 0")
	ErrRetryCountInvalid = fmt.Errorf("retry count cannot be negative")
)

type Client interface {
	ut.HealthChecker
	HTTPClient() *httpclient.Client
	GET(url string, headers http.Header) (*http.Response, error)
	POST(url string, body io.Reader, headers http.Header) (*http.Response, error)
	PUT(url string, body io.Reader, headers http.Header) (*http.Response, error)
	PATCH(url string, body io.Reader, headers http.Header) (*http.Response, error)
	DELETE(url string, headers http.Header) (*http.Response, error)
}

type ClientOptions httpclient.Option

// WithClientHTTPTimeout 设置超时时间
func WithClientHTTPTimeout(timeout time.Duration) ClientOptions {
	return ClientOptions(httpclient.WithHTTPTimeout(timeout))
}

// WithClientRetryCount 设置重试次数
func WithClientRetryCount(retryCount int) ClientOptions {
	return ClientOptions(httpclient.WithRetryCount(retryCount))
}

// WithClientRetrier 设置重试策略
func WithClientRetrier(retrier heimdall.Retriable) ClientOptions {
	return ClientOptions(httpclient.WithRetrier(retrier))
}

// WithHTTPClient 设置自定义HTTP客户端
func WithHTTPClient(client heimdall.Doer) ClientOptions {
	return ClientOptions(httpclient.WithHTTPClient(client))
}

type client struct {
	client    *httpclient.Client
	connected atomic.Bool
}

func NewClient(opts ...ClientOptions) Client {
	var c = new(client)
	var o []httpclient.Option

	for _, v := range opts {
		o = append(o, httpclient.Option(v))
	}

	c.client = httpclient.NewClient(o...)
	c.connected.Store(true)

	return c
}

func (c *client) HealthCheck(ctx context.Context) error {
	if c.client == nil {
		c.connected.Store(false)
		return fmt.Errorf("HTTP client not initialized")
	}

	c.connected.Store(true)
	return nil
}

func (c *client) IsConnected() bool {
	return c.connected.Load()
}

func (c *client) HTTPClient() *httpclient.Client {
	return c.client
}

func validateURL(urlStr string) error {
	if urlStr == "" {
		return ErrUrlEmpty
	}

	if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
		return ErrUrlInvalid
	}

	_, err := url.Parse(urlStr)
	if err != nil {
		return fmt.Errorf("URL parsing failed: %w", err)
	}

	return nil
}

func (c *client) GET(url string, headers http.Header) (*http.Response, error) {
	if err := validateURL(url); err != nil {
		return nil, err
	}
	return c.client.Get(url, headers)
}

func (c *client) POST(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	if err := validateURL(url); err != nil {
		return nil, err
	}
	return c.client.Post(url, body, headers)
}

func (c *client) PUT(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	if err := validateURL(url); err != nil {
		return nil, err
	}
	return c.client.Put(url, body, headers)
}

func (c *client) PATCH(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	if err := validateURL(url); err != nil {
		return nil, err
	}
	return c.client.Patch(url, body, headers)
}

func (c *client) DELETE(url string, headers http.Header) (*http.Response, error) {
	if err := validateURL(url); err != nil {
		return nil, err
	}
	return c.client.Delete(url, headers)
}
