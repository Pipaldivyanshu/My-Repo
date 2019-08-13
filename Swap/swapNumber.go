package main

import "fmt"

func main() {

	fmt.Println("Enter the first number")
	var first int
	fmt.Scanf("%d", &first)

	fmt.Println("Enter second number")
	var second int
	fmt.Scanf("%d", &second)

	first = first - second
	second = first + second
	first = second - first

	fmt.Println("First Number :", first)
	fmt.Println("Second Number :", second)

}
