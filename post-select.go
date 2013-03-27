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
func receive_1_1(inputs []chan int) <-chan *Output {
	cases := make([]reflect.SelectCase, len(inputs))
	output := make(chan *Output)
	doneChan := make(chan int)

	for i, input := range inputs {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(input)}
	}

	go func(numChan int) {
		defer close(output)
		for i := numChan; i > 0; i-- {
			<-doneChan
		}
	}(len(inputs))

	for {
		index, recv, ok := reflect.Select(cases)

		if !ok {
			doneChan <- 1
			break
		}

		output <- &Output{Index: index, Payload: int(recv.Int())}
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
	fmt.Println("Go 1.1:")
	printOutput(receive_1_1(getChannels()))
}
