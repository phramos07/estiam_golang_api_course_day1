package main

import "fmt"

/*Triangle area.

Write a program that calculates the area of a triangle.

The program should ask the user for input.

The user should input the BASE and HEIGHT of the triangle.

The program should output the triangle area.
*/

/*
What is a program?

A program is a value-address pair.
*/

func main() {
	var base, height float64 // the zero value

	base = 1.0   // dead code
	height = 2.0 // dead code

	// float zero value - 0.0
	// string zero value - ""
	// boolean zero value - false
	// and so on.

	fmt.Println("Enter the base:")
	_, err := fmt.Scan(&base)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Enter height:")
	fmt.Scan(height)

	area := base * height * 0.5

	fmt.Println("Area: ", area)
}
