package main

import (
	"fmt"
)

func main() {
	a := 10 //1010
	b := 3  //0011
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println(float32(a) / float32(b)) //type conversion

	//bit operators
	fmt.Println(a & b)  //0010
	fmt.Println(a | b)  //1011
	fmt.Println(a ^ b)  //1001
	fmt.Println(a &^ b) //and not .. 0100

	//bit shifting
	c := 8              //2^3
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6  shift to the left
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0 shift to the right

	// text types string: unicode8 char, runes
	s := "this is a string"
	fmt.Printf("%v, %T\n", s[2], s[2])
}
