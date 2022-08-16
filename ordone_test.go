package pipeline

import (
	"testing"
)

func TestOrDone(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		inputCh := make(chan bool, 1)
		inputCh <- true
		close(inputCh)

		neverDone := make(chan struct{})
		outputCh := OrDone(neverDone, inputCh)
		defer close(neverDone)

		val := <-outputCh
		if val != true {
			t.Errorf("Expected output to be true, got %v", val)
		}
		_, ok := <-outputCh
		if ok {
			t.Errorf("Expected outputCh to be closed, but it was not")
		}
	})

	t.Run("test done channel", func(t *testing.T) {
		inputCh := make(chan bool)
		defer close(inputCh)

		doneCh := make(chan struct{}, 1)
		doneCh <- struct{}{}
		close(doneCh)

		outputCh := OrDone(doneCh, inputCh)

		_, ok := <-outputCh
		if ok {
			t.Errorf("Expected outputCh to be closed, but it was not")
		}
	})
}
