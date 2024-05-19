package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func stage_wrapper(in In, done In, stage Stage) Out {
	out := make(Bi)
	stage_out := stage(in)
	go func() {
		defer close(out)
		for i := range stage_out {
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
	stage_in := in
	for _, stage := range stages {
		// stage_in = stage(stage_in)
		stage_in = stage_wrapper(stage_in, done, stage)
	}
	return stage_in
}
