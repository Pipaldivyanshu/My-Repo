package main

import (
	"fmt"
)

func main() {

	var n int

	a := 0
	b := 1
	newTerm := 0

	fmt.Println("Enter the integers ")
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {

		if i == 1 {

			fmt.Println(a)
		} else if i == 2 {

			fmt.Println(b)

		} else {

			newTerm = a + b
			a = b
			b = newTerm
			fmt.Println(newTerm)

		}

	}

}
