package main

import (
    "fmt"
    "math/rand"
)

type Output struct {
    Index int
    Payload interface{}
}

func receive_1(inputs interface{}) (output <-chan *Output) {
    doneChan := make(chan int)
    inputChans := inputs.([]interface{})
    outputChan := make(chan *Output)

    defer close(doneChan)

    go func(numChan int) {
        for i := numChan; i > 0; i-- {
            <-doneChan
        }
        close(outputChan)
    }(len(outputChan))

    for i, inputChan := range inputChans {
        go func() { outputChan <- &Output{ Index: i, Payload: <-inputChan.(chan interface{})} }()
    }

    return outputChan
}

func main() {
    inputs := make([]chan int, rand.Intn(10))
    fmt.Println(len(inputs))
    for i := 0; i < len(inputs); i++ {
        inputs[i] = make(chan int)
        go func() {
            for j := 0; j < rand.Intn(1000); j++ {
                inputs[i] <- j
            }
        }()
    }

    fmt.Println(len(inputs))

    output := receive_1(inputs)

    for {
        select {
        case rs, ok := <-output:
            if !ok {
                break
            }
            fmt.Printf("Received %v from %d", rs.Index, rs.Payload)
        }
    }

}
