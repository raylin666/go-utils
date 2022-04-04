package http

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	client := NewClient(
		WithClientHTTPTimeout(3 * time.Second),
		WithClientRetryCount(3))
	resp, err := client.GET("http://baidu.com", http.Header{})
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	t.Log(string(body))
}