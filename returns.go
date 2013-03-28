package main

import (
	"fmt"
	"reflect"
)

//START OMIT

func forever() {
    for {
        //do stuff
    }
    return // HL
}

func branching(x int) bool {
    if x > 10 {
        return true
    } else {
        return false
    }
    panic("This should not be happening") // HL
}

//END OMIT

func main() {
}
