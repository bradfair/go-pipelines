package pipeline

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		outputCh := RepeatFunc(ctx, func() bool {
			return true
		})

		val1 := <-outputCh
		if val1 != true {
			t.Errorf("Expected output to be true, got %v", val1)
		}
		val2 := <-outputCh
		if val2 != true {
			t.Errorf("Expected output to be true, got %v", val2)
		}
	})

	t.Run("test ctx.Done()", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		outputCh := RepeatFunc(ctx, func() bool {
			return true
		})

		ctx <- struct{}{}

		_, ok := <-outputCh
		if ok {
			t.Errorf("Expected output to be closed, but was not")
		}
	})
}
