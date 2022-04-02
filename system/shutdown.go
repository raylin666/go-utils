package system

import (
	"os"
	"os/signal"
	"syscall"
)

var _ Shutdown = (*shutdown)(nil)

// Hook a graceful shutdown hook, default with signals of SIGINT and SIGTERM
type Shutdown interface {
	// WithSignals add more signals into hook
	WithSignals(signals ...syscall.Signal) Shutdown
	// Close register shutdown handles
	Close(funcs ...func())
}

type shutdown struct {
	ctx chan os.Signal
}

// NewShutdown create a Shutdown Hook instance
func NewShutdown() Shutdown {
	hook := &shutdown{
		ctx: make(chan os.Signal, 1),
	}

	return hook.WithSignals(
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	)
}

func (h *shutdown) WithSignals(signals ...syscall.Signal) Shutdown {
	for _, s := range signals {
		signal.Notify(h.ctx, s)
	}

	return h
}

func (h *shutdown) Close(funcs ...func()) {
	select {
	case <-h.ctx:
	}
	signal.Stop(h.ctx)

	for _, f := range funcs {
		f()
	}
}
