package main

import (
  "fmt"
  "unicode/utf8"
)

//START OMIT
func main() {
    fmt.Printf("%+q\n", utf8.RuneError)
    fmt.Printf("%+q\n", string(0xD800)) //Prints utf8.RuneError
}
//END OMIT
