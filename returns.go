package main

import (
	"fmt"
	"reflect"
)

//START OMIT

func swap(in []reflect.Value) []reflect.Value {
	return []reflect.Value{in[1], in[0]}
}

func makeSwap(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()

	v := reflect.MakeFunc(fn.Type(), swap)

	fn.Set(v)
}

func main() {
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))

	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14))
}

//END OMIT
