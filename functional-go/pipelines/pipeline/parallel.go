package pipeline

import (
	"fmt"
	"strings"
)

// string transformation
type Processor func(string) string

var uppercase = func(s string) string { return strings.ToUpper(s) }
var addPrefix = func(s string) string {
	return "[PREFIX] " + s
}
var reverse = func(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func fanOutPipeline(input string, processors ...Processor) []string {
	results := make([]string, len(processors))
	done := make(chan struct {
		index int
		value string
	})

	for i, p := range processors {
		go func(i int, proc Processor) {
			output := proc(input)
			done <- struct {
				index int
				value string
			}{
				i, output,
			}
		}(i, p)
	}

	for i := 0; i < len(processors); i++ {
		res := <-done
		results[res.index] = res.value
	}
	return results
}

func RunParallelFanOut() {
	fmt.Println("\n ---[HOF : Fan-Out Pipeline] ---")

	input := "Hello dude"

	//fan-out pipeline
	results := fanOutPipeline(input, uppercase, addPrefix, reverse)

	for i, out := range results {
		fmt.Printf("-> Result #%d : %s \n", i+1, &out)
	}

}
