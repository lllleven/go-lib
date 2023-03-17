package timeoutexec

import (
	"context"
	"fmt"
	"time"
)

type FuncExecutor struct {
	f       func(...interface{}) error
	timeout time.Duration
	args    []interface{}
}

func NewFuncExecutor(f func(...interface{}) error) *FuncExecutor {
	return &FuncExecutor{
		f:       f,
		timeout: 10 * time.Second,
	}
}

func (e *FuncExecutor) WithTimeout(timeout time.Duration) *FuncExecutor {
	e.timeout = timeout
	return e
}

func (e *FuncExecutor) WithArgs(args ...interface{}) *FuncExecutor {
	e.args = args
	return e
}

func (e *FuncExecutor) Execute(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, e.timeout)
	defer cancel()

	done := make(chan error, 2)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				done <- fmt.Errorf("panic occurred: %v", r)
			}
		}()
		err := e.f(e.args...)
		done <- err
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return fmt.Errorf("function timed out")
	}
}
