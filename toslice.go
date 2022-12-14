package pipeline

import (
	"context"
)

// ToSlice converts up to the specified number of elements from the specified input channel into a slice. If the input channel is closed or the context is canceled before the specified number of elements are read, the slice will be shorter than the specified number of elements.
func ToSlice[Input any](ctx context.Context, input <-chan Input, count int) []Input {
	out := make([]Input, 0, count)
	for {
		select {
		case <-ctx.Done():
			return out
		case val, ok := <-input:
			if !ok {
				return out
			}
			out = append(out, val)
			if len(out) == count {
				return out
			}
		}
	}
}
