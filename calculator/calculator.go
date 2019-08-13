package main

import (
	"fmt"

	"./a"
	"./d"
	"./m"
	"./s"
)

func main() {

	fmt.Println("Calculator")
	fmt.Println(a.Add(6, 7))
	fmt.Println(s.Subtract(5, 9))
	fmt.Println(m.Multiply(5, 9))
	fmt.Println(d.Divide(15, 5))

}
