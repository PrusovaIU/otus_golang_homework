package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func stageWrapper(in In, done In, stage Stage) (Out, Bi) {
	out := make(Bi)
	stageOut := stage(in)
	stageDone := make(Bi)
	go func() {
		defer close(out)
		for i := range stageOut {
			select {
			case <-done:
				close(stageDone)
				return
			case out <- i:
			}
		}
	}()
	return out, stageDone
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	stageIn := in
	stageDone := done
	for _, stage := range stages {
		// stage_in = stage(stage_in)
		stageIn, stageDone = stageWrapper(stageIn, stageDone, stage)
	}
	return stageIn
}
