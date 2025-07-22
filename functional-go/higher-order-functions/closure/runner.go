package closure

import (
	"errors"
	"fmt"
	"time"
)

func BackoffRunner(base func() error) func() error {
	attempt := 0
	return func() error {
		attempt++
		delay := time.Duration(attempt) * 500 * time.Millisecond

		fmt.Printf("[BACKOFF] Attempt %d - waiting %v ... \n", attempt, delay)
		time.Sleep(delay)

		err := base()
		if err != nil {
			fmt.Printf("[BACKOFF] Error : %v \n", err)
		}
		return err
	}
}

func RunClosureState() {
	fmt.Println("\n --- [HOF : Closure State] ---")

	TaskFunc := func() error {
		fmt.Println("-> Running task. ..")
		return errors.New("still failing")
	}

	backoffTask := BackoffRunner(TaskFunc)

	for i := 0; i < 3; i++ {
		if err := backoffTask(); err != nil {
			fmt.Printf("[Run %d] Task failed: %v \n", i+1, err)
		}
	}
}
