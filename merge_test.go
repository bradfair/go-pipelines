package pipeline

import (
	"testing"
)

func TestMerge(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh1 := make(chan bool, 1)
		inputCh1 <- true
		close(inputCh1)

		inputCh2 := make(chan bool, 1)
		inputCh2 <- true
		close(inputCh2)

		outputCh := Merge(ctx, inputCh1, inputCh2)

		val1 := <-outputCh
		if val1 != true {
			t.Errorf("Expected output to be true, got %v", val1)
		}
		val2 := <-outputCh
		if val2 != true {
			t.Errorf("Expected output to be true, got %v", val2)
		}

		_, ok := <-outputCh
		if ok {
			t.Errorf("Expected output to be closed, but was not")
		}
	})

	t.Run("test ctx.Done()", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh1 := make(chan bool)
		defer close(inputCh1)

		outputCh := Merge(ctx, inputCh1)

		ctx <- struct{}{}

		_, ok := <-outputCh
		if ok {
			t.Errorf("Expected output to be closed, but was not")
		}
	})

}
