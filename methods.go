package main

import (
	"fmt"
	"strings"
)

//START OMIT
type Foo struct {
	Rotation rune
}

func (me *Foo) rot(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+me.Rotation)%26
	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+me.Rotation)%26
	}
	me.Rotation += 1
	return r
}

func main() {
	foo := &Foo{Rotation: 1}
	fmt.Println(strings.Map(foo.rot, "'Twas brillig and the slithy gopher..."))
}

//END OMIT
