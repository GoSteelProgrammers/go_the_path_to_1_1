package main

import "fmt"

//START OMIT
func main() {
    fmt.Printf("%+q\n", string(0xD800)) //Prints utf8.RuneError
}
//END OMIT
