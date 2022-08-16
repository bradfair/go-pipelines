package pipeline

import "time"

type mockContext chan struct{}

func (mockContext) Deadline() (deadline time.Time, ok bool) {
	return time.Now().Add(10 * time.Second), true
}

func (c mockContext) Done() <-chan struct{} {
	return c
}

func (mockContext) Err() error {
	return nil
}

func (mockContext) Value(key any) any {
	return nil
}
