package task

import (
	"context"
	"errors"
	"time"
)

func Timeout(d time.Duration) func(TaskFunc) TaskFunc {
	return func(fn TaskFunc) TaskFunc {
		return func() error {
			ctx, cancel := context.WithTimeout(context.Background(), d)
			defer cancel()

			done := make(chan error, 1)
			go func() {
				done <- fn()
			}()

			select {
			case <-ctx.Done():
				return errors.New("task timed out")
			case err := <-done:
				return err
			}
		}
	}
}
