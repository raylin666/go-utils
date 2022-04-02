package errors

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"runtime"
)

func callers() []uintptr {
	var pcs [32]uintptr
	l := runtime.Callers(3, pcs[:])
	return pcs[:l]
}

var _ Error = (*errorItem)(nil)
var _ fmt.Formatter = (*errorItem)(nil)

type Error interface {
	error
}

type errorItem struct {
	msg   string
	stack []uintptr
}

// New create a new error
func New(msg string) Error {
	var err = new(errorItem)
	err.msg = msg
	err.stack = callers()
	return err
}

// Errorf create a new error
func Errorf(format string, args ...interface{}) Error {
	var err = new(errorItem)
	err.msg = fmt.Sprintf(format, args...)
	err.stack = callers()
	return err
}

// Wrap with some extra message into err
func Wrap(err error, msg string) Error {
	if err == nil {
		return nil
	}

	e, ok := err.(*errorItem)
	if !ok {
		return &errorItem{msg: fmt.Sprintf("%s; %s", msg, err.Error()), stack: callers()}
	}

	e.msg = fmt.Sprintf("%s; %s", msg, e.msg)
	return e
}

// Wrapf with some extra message into err
func Wrapf(err error, format string, args ...interface{}) Error {
	if err == nil {
		return nil
	}

	msg := fmt.Sprintf(format, args...)

	e, ok := err.(*errorItem)
	if !ok {
		return &errorItem{msg: fmt.Sprintf("%s; %s", msg, err.Error()), stack: callers()}
	}

	e.msg = fmt.Sprintf("%s; %s", msg, e.msg)
	return e
}

// WithStack add caller stack information
func WithStack(err error) Error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*errorItem); ok {
		return e
	}

	return &errorItem{msg: err.Error(), stack: callers()}
}

func (e errorItem) Error() string {
	return e.msg
}

// Format used by go.uber.org/zap in Verbose
func (e errorItem) Format(f fmt.State, verb rune) {
	io.WriteString(f, e.msg)
	io.WriteString(f, "\n")
	for _, pc := range e.stack {
		fmt.Fprintf(f, "%+v\n", errors.Frame(pc))
	}
}

