package main

import "fmt"

func main() {
	var n, sum int

	fmt.Println("Enter the value of n")
	fmt.Scanf("%d", &n)

	for i := 1; i < n; i++ {

		sum = sum + i

	}
	fmt.Println("Sum of numbers is :", sum)

}
