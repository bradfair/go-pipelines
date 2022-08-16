package pipeline

import (
	"context"
	"sync"
)

// Merge merges the input channels into a single output channel.
func Merge[Input any](ctx context.Context, inputChannels ...<-chan Input) <-chan Input {
	var wg sync.WaitGroup
	multiplexedStream := make(chan Input)

	multiplex := func(ctx context.Context, inputs <-chan Input) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case input, ok := <-inputs:
				if !ok {
					return
				}
				multiplexedStream <- input
			}
		}
	}

	wg.Add(len(inputChannels))
	for _, c := range inputChannels {
		go multiplex(ctx, c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
