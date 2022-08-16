package pipeline

import (
	"testing"
)

func TestBuffer(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan int)
		defer close(inputCh)

		outputCh := Buffer(ctx, inputCh, 2)

		inputCh <- 1
		inputCh <- 2

		val := <-outputCh
		if val != 1 {
			t.Errorf("Expected output to be 1, got %d", val)
		}
		val = <-outputCh
		if val != 2 {
			t.Errorf("Expected output to be 2, got %d", val)
		}
	})

	t.Run("test ctx.Done()", func(t *testing.T) {
		ctx := make(mockContext)

		inputCh := make(chan int)
		defer close(inputCh)

		outputCh := Buffer(ctx, inputCh, 2)

		ctx <- struct{}{}

		_, ok := <-outputCh
		if ok {
			t.Errorf("Expected output to be closed, but was not")
		}
	})
}
