package pipeline

import (
	"context"
	"fmt"
	"testing"
)

func TestToSlice(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan bool, 1)
		inputCh <- true
		close(inputCh)

		output := ToSlice(ctx, inputCh, 1)

		if fmt.Sprintf("%v", output) != "[true]" {
			t.Error("Expected output to be []bool.")
		}
	})

	t.Run("test ctx.Done()", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		inputCh := make(chan bool)
		defer close(inputCh)

		output := ToSlice(ctx, inputCh, 1)

		if fmt.Sprintf("%v", output) != "[]" {
			t.Error("Expected output to be empty due to ctx.Done().")
		}
	})
}
