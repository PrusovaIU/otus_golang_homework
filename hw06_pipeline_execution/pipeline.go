package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func chanalWrapper(in In, done In) Bi {
	chIn := make(Bi)
	go func() {
		defer close(chIn)
		for {
			select {
			case <-done:
				return
			case value, ok := <-in:
				if !ok {
					return
				}
				chIn <- value
			}
		}
	}()
	return chIn
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	var out Out
	for _, stage := range stages {
		in = chanalWrapper(in, done)
		out = stage(in)
		in = out
	}
	return chanalWrapper(out, done)
}
