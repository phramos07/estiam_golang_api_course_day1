package main

import "fmt"

/*
Write a Go program that takes a student's score as input and then outputs a corresponding letter grade based on the following scale:

90-100: A
80-89: B
70-79: C
60-69: D
0-59: F
*/

// two solutions - one with IF/ELSE and one with SWITCH/CASE
func main() {
	var score int

	fmt.Println("Enter your score: ")
	fmt.Scan(&score)

	var grade string
	if score >= 90 && score <= 100 {
		grade = "A"
	} else if score >= 80 && score < 90 {
		grade = "B"
	} else if score >= 70 && score < 80 {
		grade = "C"
	} else if score >= 60 && score < 70 {
		grade = "D"
	} else {
		grade = "F"
	}

	// Using a switch
	switch {
	case score >= 90 && score <= 100:
		grade = "A"
	case score >= 80 && score <= 90:
		grade = "B"
	case score >= 70 && score <= 80:
		grade = "C"
	case score >= 60 && score <= 70:
		grade = "D"
	default:
		grade = "F"
	}

	fmt.Printf("Grade: %s", grade)
}
