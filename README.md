# pipeline

[![codecov](https://codecov.io/gh/bradfair/go-pipelines/branch/main/graph/badge.svg?token=1DI2ONSVP4)](https://codecov.io/gh/bradfair/go-pipelines)

## Functions

### func [AggregatorFunc](/aggregator.go#L8)

`func AggregatorFunc[Input, Output any](ctx context.Context, input <-chan Input, count int, f func(ctx context.Context, input ...Input) Output) <-chan Output`

AggregatorFunc uses the given function to aggregate n Inputs to one Output.

### func [Bridge](/bridge.go#L8)

`func Bridge[Input any](ctx context.Context, inputsStream <-chan <-chan Input) <-chan Input`

Bridge takes a channel of input channels and combines their elements into a single output channel.

### func [Buffer](/buffer.go#L8)

`func Buffer[Input any](ctx context.Context, input <-chan Input, count int) <-chan Input`

Buffer converts a given channel into a buffered channel.

### func [DecomposerFunc](/decomposer.go#L8)

`func DecomposerFunc[Input, Output any](ctx context.Context, input <-chan Input, f func(input Input) []Output) <-chan Output`

DecomposerFunc uses the given function to decompose each input from a channel of generic Inputs to a channel of generic Outputs.

### func [FilterFunc](/filter.go#L8)

`func FilterFunc[Input any](ctx context.Context, input <-chan Input, f func(ctx context.Context, input Input) bool) (_, _ <-chan Input)`

FilterFunc takes a context, input channel, and a filter function that accepts an input and returns a boolean. It returns two channels. One that blocks with the inputs that pass the filter, and one doesn't block with the inputs that do not pass the filter.

### func [Merge](/merge.go#L9)

`func Merge[Input any](ctx context.Context, inputChannels ...<-chan Input) <-chan Input`

Merge merges the input channels into a single output channel.

### func [OrDone](/ordone.go#L4)

`func OrDone[Input any](done <-chan struct{}, inputs <-chan Input) <-chan Input`

OrDone wraps a channel with a done channel and returns a forwarding channel that closes when either the original channel or the done channel closes.

### func [RepeatFunc](/repeat.go#L7)

`func RepeatFunc[Output any](ctx context.Context, fn func() Output) <-chan Output`

### func [SinkFunc](/sink.go#L8)

`func SinkFunc[Input any](ctx context.Context, inputs <-chan Input, f func(ctx context.Context, input Input))`

SinkFunc executes a function on each item in a channel until the channel is closed or the context is cancelled.

### func [Take](/take.go#L8)

`func Take[Input any](ctx context.Context, input <-chan Input, count int) <-chan Input`

Take returns a channel that closes after receiving the specified number of elements from the specified input channel.

### func [Tee](/tee.go#L8)

`func Tee[Input any](ctx context.Context, input <-chan Input) (_, _ <-chan Input)`

Tee splits the input channel into two output channels, and blocks reading the next input until both output channels receive each element.

### func [ToSlice](/toslice.go#L8)

`func ToSlice[Input any](ctx context.Context, input <-chan Input, count int) []Input`

ToSlice converts up to the specified number of elements from the specified input channel into a slice. If the input channel is closed or the context is canceled before the specified number of elements are read, the slice will be shorter than the specified number of elements.

### func [TransformerFunc](/transformer.go#L8)

`func TransformerFunc[Input, Output any](ctx context.Context, input <-chan Input, f func(ctx context.Context, input Input) Output) <-chan Output`

TransformerFunc uses the given function to transform a channel of generic Inputs to a channel of generic Outputs.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
