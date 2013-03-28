package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type Output struct {
	Index   int
	Payload int
}

func receive(inputs []chan int) <-chan *Output {
	//START OMIT
	cases := make([]reflect.SelectCase, len(inputs))
	output := make(chan *Output)

	for i, input := range inputs {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(input)}
	}

	go func() {
		defer close(output)
		for {
			index, recv, ok := reflect.Select(cases)
			if !ok {
				cases[index] = cases[len(cases)-1]
				cases = cases[0 : len(cases)-1]

				if len(cases) == 0 {
					break
				}
				continue
			}

			output <- &Output{Index: index, Payload: int(recv.Int())}
		}
	}()
	//END OMIT

	return output
}

func getChannels() []chan int {
	inputs := make([]chan int, 10)
	for i := 0; i < len(inputs); i++ {
		input := make(chan int)
		go func() {
			defer close(input)
			for j := 0; j < 10; j++ {
				input <- j
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
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
	printOutput(receive(getChannels()))
}
