package pipeline

import (
	"context"
)

// Tee splits the input channel into two output channels, and blocks reading the next input until both output channels receive each element.
func Tee[Input any](ctx context.Context, input <-chan Input) (_, _ <-chan Input) {
	out1 := make(chan Input)
	out2 := make(chan Input)
	go func(ctx context.Context, input <-chan Input, out1, out2 chan<- Input) {
		defer close(out1)
		defer close(out2)
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-input:
				if !ok {
					return
				}
				var out1, out2 = out1, out2
				for i := 0; i < 2; i++ {
					select {
					case out1 <- val:
						out1 = nil
					case out2 <- val:
						out2 = nil
					}
				}
			}
		}
	}(ctx, input, out1, out2)
	return out1, out2
}
