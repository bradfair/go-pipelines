package pipeline

import (
	"context"
)

// Bridge takes a channel of input channels and combines their elements into a single output channel.
func Bridge[Input any](ctx context.Context, inputsStream <-chan <-chan Input) <-chan Input {
	out := make(chan Input)
	go receiveInputStreams(ctx, inputsStream, out)
	return out
}

func receiveInputStreams[Input any](ctx context.Context, inputsStream <-chan <-chan Input, out chan Input) {
	defer func() {
		select {
		case <-out:
		default:
			close(out)
		}
	}()
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return
		case inputStream, ok := <-inputsStream:
			if !ok {
				return
			}
			if ctx.Err() != nil {
				return
			}
			go bridgeInputStream(childCtx, inputStream, out)

		}
	}
}

func bridgeInputStream[Input any](ctx context.Context, inputs <-chan Input, out chan<- Input) {
	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-inputs:
			if !ok {
				return
			}
			if ctx.Err() != nil {
				return
			}
			out <- val
		}
	}
}
