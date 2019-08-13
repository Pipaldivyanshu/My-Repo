package main

import "fmt"

func main() {

	var r float64
	var area float64
	var circ float64
	const PI float64 = 3.14

	fmt.Println("Enter the radius")
	fmt.Scanf("%f", &r)

	area = PI * r * r
	fmt.Println("Area of circle is :", area)

	circ = 2 * PI * r
	fmt.Println("Circumference of circle is :", circ)

}
