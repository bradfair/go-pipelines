package pipeline

import (
	"context"
)

// SinkFunc executes a function on each item in a channel until the channel is closed or the context is cancelled.
func SinkFunc[Input any](ctx context.Context, inputs <-chan Input, f func(ctx context.Context, input Input)) {
	go func() {
		for input := range OrDone(ctx.Done(), inputs) {
			f(ctx, input)
		}
	}()
}
