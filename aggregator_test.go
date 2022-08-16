package pipeline

import (
	"context"
	"testing"
)

func TestAggregatorFunc(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan bool, 2)
		inputCh <- true
		inputCh <- true
		close(inputCh)

		outputCh := AggregatorFunc(ctx, inputCh, 2, func(ctx context.Context, input ...bool) int {
			count := 0
			for _, val := range input {
				if val {
					count++
				}
			}
			return count
		})
		val := <-outputCh
		if val != 2 {
			t.Errorf("Expected output to be 2, got %d", val)
		}
		_, ok := <-outputCh
		if ok {
			t.Errorf("Expected output to be closed, but was not")
		}
	})

	t.Run("test done before work starts", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan bool)
		defer close(inputCh)

		outputCh := AggregatorFunc(ctx, inputCh, 2, func(ctx context.Context, input ...bool) int {
			count := 0
			for _, val := range input {
				if val {
					count++
				}
			}
			return count
		})

		ctx <- struct{}{}

		_, ok := <-outputCh
		if ok {
			t.Errorf("Expected output to be closed, got %v", ok)
		}
	})

	t.Run("test done during work", func(t *testing.T) {
		ctx := make(mockContext)
		defer close(ctx)

		inputCh := make(chan bool)
		defer close(inputCh)

		outputCh := AggregatorFunc(ctx, inputCh, 2, func(ctx context.Context, input ...bool) int {
			count := 0
			for _, val := range input {
				if val {
					count++
				}
			}
			return count
		})

		inputCh <- true
		inputCh <- true

		ctx <- struct{}{}

		val := <-outputCh
		if val != 2 {
			t.Errorf("Expected output to be 2, got %d", val)
		}
		_, ok := <-outputCh
		if ok {
			t.Errorf("Expected output to be closed, got %v", ok)
		}
	})
}
