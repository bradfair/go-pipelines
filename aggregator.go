package pipeline

import (
	"context"
)

// AggregatorFunc uses the given function to aggregate n Inputs to one Output.
func AggregatorFunc[Input, Output any](ctx context.Context, input <-chan Input, count int, f func(ctx context.Context, input ...Input) Output) <-chan Output {
	out := make(chan Output)
	go func(ctx context.Context, input <-chan Input, out chan<- Output, count int, f func(ctx context.Context, input ...Input) Output) {
		defer close(out)
		var inputs []Input
		dump := func() {
			if len(inputs) > 0 {
				out <- f(ctx, inputs...)
			}
		}
		for {
			select {
			case <-ctx.Done():
				dump()
				return
			case val, ok := <-input:
				if !ok {
					dump()
					return
				}
				inputs = append(inputs, val)
				if len(inputs) == count {
					output := f(ctx, inputs...)
					select {
					case <-ctx.Done():
						dump()
						return
					case out <- output:
						inputs = nil
					}
				}
			}
		}
	}(ctx, input, out, count, f)
	return out
}
