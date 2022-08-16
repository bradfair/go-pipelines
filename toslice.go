package pipeline

import (
	"context"
)

// ToSlice converts the specified number of elements from the specified input channel into a slice.
func ToSlice[Input any](ctx context.Context, input <-chan Input, count int) []Input {
	var out []Input
	for {
		select {
		case <-ctx.Done():
			return out
		case val := <-input:
			out = append(out, val)
			if len(out) == count {
				return out
			}
		}
	}
}
