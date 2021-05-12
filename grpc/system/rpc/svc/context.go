package svc

import (
	"fmt"
	"github.com/raylin666/go-gin-api/context"
)

type Context struct {
	*context.Context
}

func NewContext() *Context {
	fmt.Println(1)
	return &Context{}
}