package main

import (
	"fmt"
	"reflect"
)

type Output struct {
	Index   int
	Payload int
}

//START OMIT

//TODO Replace done chan with waitgroup
func receive_1(inputs []chan int) <-chan *Output {
	doneChan := make(chan int)
	output := make(chan *Output)

	go func(numChan int) {
		defer close(output)
		for i := numChan; i > 0; i-- {
			<-doneChan
		}
	}(len(inputs))

	for i, input := range inputs {
		go func(index int) {
		L:
			for {
				select {
				case payload, ok := <-input:
					if !ok {
						doneChan <- 1
						break L
					}
					output <- &Output{Index: index, Payload: payload}
				}
			}
		}(i)
	}

	return output
}

//END OMIT

func getChannels() []chan int {
	inputs := make([]chan int, 1000)
	for i := 0; i < len(inputs); i++ {
		input := make(chan int)
		go func() {
			defer close(input)
			for j := 0; j < 1000; j++ {
				input <- j
			}
		}()
		inputs[i] = input
	}
	return inputs
}

func printOutput(output <-chan *Output) {
L:
	for {
		select {
		case rs, ok := <-output:
			if !ok {
				break L
			}
			fmt.Printf("Received from %d: %d\n", rs.Index, rs.Payload)
		}
	}
}

func main() {
	fmt.Println("Go 1.0:")
	printOutput(receive_1(getChannels()))
}
