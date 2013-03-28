package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Output struct {
	Index   int
	Payload int
}

func receive(inputs []chan int) <-chan *Output {
	//START OMIT
	var wg sync.WaitGroup
	output := make(chan *Output)

	for i, input := range inputs {
		wg.Add(1)
		go func(index int, input chan int) {
		L:
			for {
				select {
				case payload, ok := <-input:
					if !ok {
						wg.Done()
						break L
					}
					output <- &Output{Index: index, Payload: payload}
				}
			}
		}(i, input)
	}
	go func(numChan int) {
		defer close(output)
		wg.Wait()
	}(len(inputs))
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
	rand.Seed(time.Now().Unix())
	printOutput(receive(getChannels()))
}
