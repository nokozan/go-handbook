package task

func Hooks(before, after func()) func(TaskFunc) TaskFunc {
	return func(fn TaskFunc) TaskFunc {
		return func() error {
			if before != nil {
				before()
			}
			err := fn()
			if after != nil {
				after()
			}
			return err
		}
	}
}
