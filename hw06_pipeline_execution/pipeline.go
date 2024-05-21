package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func stageWrapper(in In, done In, stage Stage) Out {
	out := make(Bi)
	stageOut := stage(in)
	go func() {
		defer close(out)
		for i := range stageOut {
			select {
			case <-done:
				return
			case out <- i:
			}
		}
	}()
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	stageIn := in
	for _, stage := range stages {
		// stage_in = stage(stage_in)
		stageIn = stageWrapper(stageIn, done, stage)
	}
	return stageIn
}
