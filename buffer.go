package pipeline

import (
	"context"
)

// Buffer converts a given channel into a buffered channel.
func Buffer[Input any](ctx context.Context, input <-chan Input, count int) <-chan Input {
	out := make(chan Input, count)
	go func(ctx context.Context, input <-chan Input, out chan<- Input) {
		defer close(out)
		for i := range OrDone(ctx.Done(), input) {
			out <- i
		}
	}(ctx, input, out)
	return out
}
