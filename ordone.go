package pipeline

// OrDone wraps a channel with a done channel and returns a forwarding channel that closes when either the original channel or the done channel closes.
func OrDone[Input any](done <-chan struct{}, inputs <-chan Input) <-chan Input {
	output := make(chan Input)
	go func(done <-chan struct{}) {
		defer close(output)
		for {
			select {
			case <-done:
				return
			case input, ok := <-inputs:
				if !ok {
					return
				}
				select {
				case output <- input:
				case <-done:
				}
			}
		}
	}(done)
	return output
}
