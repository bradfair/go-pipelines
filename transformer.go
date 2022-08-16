package pipeline

import (
	"context"
)

// TransformerFunc uses the given function to transform a channel of generic Inputs to a channel of generic Outputs.
func TransformerFunc[Input, Output any](ctx context.Context, input <-chan Input, f func(ctx context.Context, input Input) Output) <-chan Output {
	out := make(chan Output)
	go func(ctx context.Context, input <-chan Input, out chan<- Output) {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case val := <-input:
				out <- f(ctx, val)
			}
		}
	}(ctx, input, out)
	return out
}
