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
		defer close(stageDone)
		for i := range stageOut {
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}()
	return out, stageDone
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	stageIn := in
	stageDone := done
	for _, stage := range stages {
		stageIn, stageDone = stageWrapper(stageIn, stageDone, stage)
	}
	return stageIn
}
