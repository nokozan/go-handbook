package task

import "fmt"

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
