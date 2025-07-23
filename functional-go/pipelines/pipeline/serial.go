package pipeline

import "fmt"

func serialPipeline(input string, processors ...Processor) string {
	output := input
	for _, fn := range processors {
		output = fn(output)
	}
	return output

}

func RunSerialPipeline() {
	input := "hello dude"

	result := serialPipeline(input, uppercase, addPrefix, reverse)
	fmt.Println("result ", result)
}
