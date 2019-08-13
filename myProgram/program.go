package main

import "fmt"

func main() {

	fmt.Println("Enter a new number")
	var n int
	fmt.Scanf("%d", &n)

	if n%2 == 0 {

		fmt.Println("Number is even")

	} else {
		fmt.Println("Number is odd")

	}
}
