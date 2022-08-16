package pipeline

import (
	"testing"
)

func TestDecomposer(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan []bool, 1)
		inputCh <- []bool{true, false}
		close(inputCh)

		outputCh := DecomposerFunc(ctx, inputCh, func(b []bool) []bool {
			var flipped []bool
			for _, v := range b {
				flipped = append(flipped, !v)
			}
			return flipped
		})

		val := <-outputCh
		if val {
			t.Errorf("Expected output to be false, got true")
		}
		val = <-outputCh
		if !val {
			t.Errorf("Expected output to be true, got false")
		}
		_, ok := <-outputCh
		if ok {
			t.Errorf("Expected output to be closed, but was not")
		}
	})

	t.Run("test ctx.Done()", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan []bool)
		defer close(inputCh)

		outputCh := DecomposerFunc(ctx, inputCh, func(b []bool) []bool {
			return b
		})

		ctx <- struct{}{}

		_, ok := <-outputCh
		if ok {
			t.Errorf("Expected output to be closed, but was not")
		}
	})
}
