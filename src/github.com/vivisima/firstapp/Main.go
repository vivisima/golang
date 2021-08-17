package main

import (
	"fmt"
)

const b = iota

func main() {
	// CONSTANTS, can be shadowed
	const myConst int = 42 // MyConst --> this will be automatically exported
	var b int = 27
	fmt.Printf("%v, %T\n", myConst+b, myConst+b)
}
