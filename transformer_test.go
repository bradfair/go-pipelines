package pipeline

import (
	"context"
	"strconv"
	"testing"
)

func TestTransformerFunc(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		ctx := make(mockContext)
		inputCh := make(chan string, 1)
		inputCh <- "1"
		close(inputCh)

		outputCh := TransformerFunc(ctx, inputCh, func(ctx context.Context, input string) int {
			if output, err := strconv.Atoi(input); err == nil {
				return output
			}
			return 0
		})

		output := <-outputCh
		if output != 1 {
			t.Errorf("Expected output to be 1, got %d", output)
		}
	})

	t.Run("test ctx.Done()", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan string)
		defer close(inputCh)

		outputCh := TransformerFunc(ctx, inputCh, func(ctx context.Context, input string) int {
			if output, err := strconv.Atoi(input); err == nil {
				return output
			}
			return 0
		})

		ctx <- struct{}{}

		_, ok := <-outputCh
		if ok {
			t.Error("Expected output channel to be closed")
		}
	})
}
