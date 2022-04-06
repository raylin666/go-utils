package http

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	s := NewServer(&http.Server{}, WithServerAddress(":10001"))
	// 处理函数
	http.Handle("/", &TServerHandler{})
	err := s.Start(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

type TServerHandler struct {}

func (t *TServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(fmt.Sprintf("hello world! %v - %v", w, r))
}