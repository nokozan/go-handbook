package task

import (
	"errors"
	"fmt"
	"time"
)

// a unit of work
type TaskFunc func() error

// E.g Executor{ Name: "SendEmail", Task: decoratedFunc }.
type Executor struct {
	Name string
	Task TaskFunc
}

func (e *Executor) Execute() error {
	fmt.Printf("[EXECUTOR] Running task: %s\n", e.Name)
	return e.Task()
}

func RunTaskWrappersExample() {
	fmt.Println("=== [HOF: Task Wrappers] ===")

	//Base task : always fails
	failingTask := func() error {
		fmt.Println("→ Executing task (simulate failure)...")
		return errors.New("simulated error")
	}

	//Compose wrappers using HOF

	decorated := Hooks(
		func() { fmt.Println("[HOOK] before execution") },
		func() { fmt.Println("[HOOK] after execution") },
	)(
		Timeout(2 * time.Second)(
			Retry(2, 500*time.Millisecond)(
				failingTask,
			),
		),
	)

	executor := Executor{
		Name: "Job",
		Task: decorated,
	}
	if err := executor.Execute(); err != nil {
		fmt.Println("[FAILURE]", err)
	}
}
