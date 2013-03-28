package main

import (
	"io"
	"os"
)

func main() {
	//START OMIT
	var w = os.Stdout
	var f1 func(p []byte) (n int, err error)
	var f2 func(w io.Writer, p []byte) (n int, err error)

	//Method value (new)
	f1 = w.Write
	f1([]byte("hello ")) //writes the byte array to w // HL

	//Method expression (pre-existing)
	f2 = (io.Writer).Write
	f2(w, []byte("world")) //writes the byte array to w // HL

	w.Write([]byte("!\n")) // HL
	//END OMIT
}
