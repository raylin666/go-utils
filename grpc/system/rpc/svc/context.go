package svc

import "github.com/raylin666/go-gin-api/context"

type Context struct {
	*context.Context
}

func NewContext() *Context {
	return &Context{}
}