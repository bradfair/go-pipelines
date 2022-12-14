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

	t.Run("test fewer inputs than count", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan bool, 1)
		inputCh <- true
		close(inputCh)

		output := ToSlice(ctx, inputCh, 10)

		if len(output) != 1 {
			t.Errorf("Expected output to be length 1, got %d", len(output))
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
