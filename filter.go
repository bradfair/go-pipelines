package pipeline

import (
	"context"
)

// FilterFunc takes a context, input channel, and a filter function that accepts an input and returns a boolean. It returns two channels. One that blocks with the inputs that pass the filter, and one doesn't block with the inputs that do not pass the filter.
func FilterFunc[Input any](ctx context.Context, input <-chan Input, f func(ctx context.Context, input Input) bool) (_, _ <-chan Input) {
	kept := make(chan Input)
	discarded := make(chan Input)
	go func(ctx context.Context, input <-chan Input, kept, discarded chan Input, f func(ctx context.Context, input Input) bool) {
		defer close(kept)
		defer close(discarded)
		for val := range OrDone(ctx.Done(), input) {
			if f(ctx, val) {
				select {
				case <-ctx.Done():
				case kept <- val:
				}
			} else {
				select {
				case <-ctx.Done():
				case discarded <- val:
				default:
				}
			}
		}
	}(ctx, input, kept, discarded, f)
	return kept, discarded
}
