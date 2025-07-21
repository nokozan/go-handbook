package task

import (
	"fmt"
	"time"
)

func Retry(times int, delay time.Duration) func(TaskFunc) TaskFunc {
	return func(fn TaskFunc) TaskFunc {
		return func() error {
			var err error
			for i := 0; i <= times; i++ {
				err = fn()
				if err == nil {
					return nil
				}
				time.Sleep(delay)
			}
			return fmt.Errorf("task failed after %d retries :%w", times, err)
		}
	}
}
