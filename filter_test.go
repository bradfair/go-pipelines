package pipeline

import (
	"context"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan bool, 2)
		inputCh <- true
		inputCh <- false
		close(inputCh)

		trueCh, falseCh := FilterFunc(ctx, inputCh, func(ctx context.Context, b bool) bool {
			return b
		})

		val := <-trueCh
		if val != true {
			t.Errorf("Expected output to be true, got false")
		}

		val = <-falseCh
		if val != false {
			t.Errorf("Expected output to be false, got true")
		}

		_, ok := <-trueCh
		if ok {
			t.Errorf("Expected output to be closed, but was not")
		}

		_, ok = <-falseCh
		if ok {
			t.Errorf("Expected output to be closed, but was not")
		}
	})

	t.Run("test ctx.Done()", func(t *testing.T) {
		ctx := make(mockContext)

		inputCh := make(chan bool)
		defer close(inputCh)

		trueCh, falseCh := FilterFunc(ctx, inputCh, func(ctx context.Context, b bool) bool {
			return b
		})

		ctx <- struct{}{}

		_, ok := <-trueCh
		if ok {
			t.Errorf("Expected output to be closed, but was not")
		}

		_, ok = <-falseCh
		if ok {
			t.Errorf("Expected output to be closed, but was not")
		}
	})
}
