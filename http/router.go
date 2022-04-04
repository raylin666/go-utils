package http

import "net/http"

type Router interface {
	ServeHTTP(writer http.ResponseWriter, request *http.Request)
}