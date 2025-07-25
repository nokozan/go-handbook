package pipeline

import (
	"fmt"

	"sync"
)

func FanInPipeline(input string, processors ...Processor) []string {
	var wg sync.WaitGroup
	outputChan := make(chan string, len(processors))

	for _, proc := range processors {
		wg.Add(1)
		go func(p Processor) {
			defer wg.Done()
			result := p(input)
			outputChan <- result
		}(proc)
	}

	go func() {
		wg.Wait()
		close(outputChan)
	}()

	var results []string
	for out := range outputChan {
		results = append(results, out)
	}

	return results

}

func RunFanInPipelineExample() {
	fmt.Println(" Fan-In Pipeline")

	procs := []Processor{
		uppercase,
		addPrefix,
		reverse,
	}

	results := FanInPipeline("hello", procs...)

	for i, r := range results {
		fmt.Printf("Result [%d] : %s\n", i, r)
	}
}
