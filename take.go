package pipeline

import (
	"context"
)

// Take returns a channel that closes after receiving the specified number of elements from the specified input channel.
func Take[Input any](ctx context.Context, input <-chan Input, count int) <-chan Input {
	output := make(chan Input)
	go func(ctx context.Context) {
		defer close(output)
		for i := 0; i < count; i++ {
			select {
			case <-ctx.Done():
				return
			case val := <-input:
				output <- val
			}
		}
	}(ctx)
	return output
}
