package main

import (
	"fmt"
)

func main() {

	x := []int{}
	x = append(x, 1)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 2)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
