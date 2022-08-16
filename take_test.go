package pipeline

import (
	"testing"
)

func TestTake(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan bool, 2)
		inputCh <- true
		inputCh <- true
		defer close(inputCh)

		outputCh := Take(ctx, inputCh, 1)

		val := <-outputCh
		if val != true {
			t.Errorf("Expected output to be true, got %v", val)
		}
		val, ok := <-outputCh
		if ok {
			t.Errorf("Expected output to be closed, got %v", val)
		}
	})

	t.Run("test ctx.Done()", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan bool)
		defer close(inputCh)

		outputCh := Take(ctx, inputCh, 100)

		ctx <- struct{}{}

		_, ok := <-outputCh
		if ok {
			t.Errorf("Expected output to be closed, got %v", ok)
		}
	})
}
