package main

import (
	"fmt"
)

func main() {

	/* files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	} */
	//var i int
	//i = 42
	//var j int = 27
	var l bool //all initialized variables have a 0 value
	k := 99
	fmt.Printf("%v, %T\n", k, k)
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", l, l)
}
