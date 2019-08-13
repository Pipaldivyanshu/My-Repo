package main

import "fmt"

func main() {

	var n int

	fmt.Println("Enter the number")
	fmt.Scanf("%d", &n)

	fact := 1

	for i := 1; i <= n; i++ {

		fact = fact * i

	}

	fmt.Printf("Factorial of %v is %v", n, fact)

}
