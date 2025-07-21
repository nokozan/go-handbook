package main

import (
	"errors"
	"fmt"
	"higher-order-functions/task"
	"time"
)

func main() {
	runTaskWrappersExample()
}

func runTaskWrappersExample() {
	fmt.Println("=== [HOF: Task Wrappers] ===")

	//Base task : always fails
	failingTask := func() error {
		fmt.Println("→ Executing task (simulate failure)...")
		return errors.New("simulated error")
	}

	//Compose wrappers using HOF

	decorated := task.Hooks(
		func() { fmt.Println("[HOOK] before execution") },
		func() { fmt.Println("[HOOK] after execution") },
	)(
		task.Timeout(2 * time.Second)(
			task.Retry(2, 500*time.Millisecond)(
				failingTask,
			),
		),
	)

	executor := task.Executor{
		Name: "Job",
		Task: decorated,
	}
	if err := executor.Execute(); err != nil {
		fmt.Println("[FAILURE]", err)
	}
}
