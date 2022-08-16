package pipeline

import (
	"testing"
)

func TestTee(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan bool, 1)
		inputCh <- true
		close(inputCh)

		outputCh1, outputCh2 := Tee(ctx, inputCh)

		got1, got2 := false, false
		for {
			select {
			case <-ctx.Done():
				t.Error("Should not have timed out")
				return
			case output := <-outputCh1:
				if output != true {
					t.Errorf("Expected output to be true, got %v", output)
				} else {
					got1 = true
				}
			case output := <-outputCh2:
				if output != true {
					t.Errorf("Expected output to be true, got %v", output)
				} else {
					got2 = true
				}
			default:
				if got1 && got2 {
					return
				}
			}
		}
	})

	t.Run("test input stream closing", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan bool)

		outputCh1, outputCh2 := Tee(ctx, inputCh)

		close(inputCh)

		_, ok := <-outputCh1
		if ok {
			t.Errorf("Expected output to be closed, got %v", ok)
		}

		_, ok = <-outputCh2
		if ok {
			t.Errorf("Expected output to be closed, got %v", ok)
		}
	})

	t.Run("test ctx.Done()", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan bool)
		defer close(inputCh)

		outputCh1, outputCh2 := Tee(ctx, inputCh)

		ctx <- struct{}{}

		_, ok := <-outputCh1
		if ok {
			t.Errorf("Expected output to be closed, got %v", ok)
		}

		_, ok = <-outputCh2
		if ok {
			t.Errorf("Expected output to be closed, got %v", ok)
		}
	})
}
