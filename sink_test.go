package pipeline

import (
	"context"
	"testing"
)

func TestSinkFunc(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan bool, 1)
		inputCh <- true
		close(inputCh)

		i := 0

		done := make(chan struct{})

		SinkFunc(ctx, inputCh, func(ctx context.Context, input bool) {
			i++
			done <- struct{}{}
		})

		<-done
		if i != 1 {
			t.Errorf("Expected i to be 1, got %d", i)
		}

	})

	t.Run("test ctx.Done()", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		var outputs []bool

		inputCh := make(chan bool, 1)
		defer close(inputCh)

		SinkFunc(ctx, inputCh, func(ctx context.Context, val bool) {
			outputs = append(outputs, val)
		})

		ctx <- struct{}{}
		inputCh <- true

		if len(outputs) != 0 {
			t.Errorf("Expected 0 outputs due to closed context, got %d", len(outputs))
		}
	})
}
