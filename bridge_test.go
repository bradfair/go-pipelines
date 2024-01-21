package pipeline

import (
	"context"
	"testing"
)

func TestBridge(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		inputChs := make(chan (<-chan bool), 2)

		inputCh1 := make(chan bool, 1)
		inputCh1 <- true
		close(inputCh1)

		inputCh2 := make(chan bool, 1)
		inputCh2 <- true
		close(inputCh2)

		inputChs <- inputCh1
		inputChs <- inputCh2

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		outputCh := Bridge(ctx, inputChs)

		val1 := <-outputCh
		if val1 != true {
			t.Errorf("Expected output to be true, got %v", val1)
		}
		val2 := <-outputCh
		if val2 != true {
			t.Errorf("Expected output to be true, got %v", val2)
		}
	})

	t.Run("test input stream closing", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		inputChs := make(chan (<-chan bool))

		outputCh := Bridge(ctx, inputChs)

		close(inputChs)

		_, ok := <-outputCh
		if ok {
			t.Errorf("Expected output channel to be closed")
		}
	})

	t.Run("test an input stream closing", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		inputChs := make(chan (<-chan bool), 2)

		inputCh1 := make(chan bool, 1)
		inputCh2 := make(chan bool)
		inputCh3 := make(chan bool, 1)

		outputCh := Bridge(ctx, inputChs)

		inputChs <- inputCh1
		inputChs <- inputCh2
		inputChs <- inputCh3

		close(inputCh2)

		inputCh1 <- true
		inputCh3 <- true

		val1 := <-outputCh
		if val1 != true {
			t.Errorf("Expected output to be true, got %v", val1)
		}
		val2 := <-outputCh
		if val2 != true {
			t.Errorf("Expected output to be true, got %v", val2)
		}
		select {
		case _, ok := <-outputCh:
			if !ok {
				t.Errorf("Expected output channel to still be open")
			}
		default:
			inputCh1 <- true
			inputCh3 <- true
		}

		_, ok := <-outputCh
		if !ok {
			t.Errorf("Expected output channel to be open")
		}

	})

	t.Run("test ctx.Done()", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		inputChs := make(chan (<-chan bool), 2)

		inputCh1 := make(chan bool, 1)
		defer close(inputCh1)

		inputCh2 := make(chan bool, 1)
		defer close(inputCh2)

		inputChs <- inputCh1
		inputChs <- inputCh2

		outputCh := Bridge(ctx, inputChs)

		cancel()

		_, ok := <-outputCh
		if ok {
			t.Errorf("Expected output to be closed, but was not")
		}
	})
}
