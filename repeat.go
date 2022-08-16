package pipeline

import (
	"context"
)

func RepeatFunc[Output any](ctx context.Context, fn func() Output) <-chan Output {
	out := make(chan Output)
	go func(ctx context.Context) {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case out <- fn():
			}
		}
	}(ctx)
	return out
}
