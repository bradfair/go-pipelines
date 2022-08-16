package pipeline

import (
	"context"
)

// DecomposerFunc uses the given function to decompose each input from a channel of generic Inputs to a channel of generic Outputs.
func DecomposerFunc[Input, Output any](ctx context.Context, input <-chan Input, f func(input Input) []Output) <-chan Output {
	out := make(chan Output)
	go func(ctx context.Context, input <-chan Input, out chan<- Output) {
		defer close(out)
		for i := range OrDone(ctx.Done(), input) {
			outputs := f(i)
			for _, output := range outputs {
				out <- output
			}
		}
	}(ctx, input, out)
	return out
}
