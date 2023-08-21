package main

import (
	"fmt"
	"math"
)

/*Write a Go program that takes an integer input from the user and checks whether the entered number is a prime number or not.

How do I check if a number is a prime or not?

1. get the square root of the number sqrt(x)
2. round the square root of that number to integer
3. for every number from 2.... to int(sqrt(x)) ... If it divides the number, then it is not prime.

11 is prime
17 is prime
13 is prime
21 is not
33 is not
37 is prime
42 is not
47 is prime
50 is not
51 is not
*/

func main() {
	var number int

	fmt.Println("Enter the number:")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println(err)
		return
	}

	if number == 1 {
		fmt.Println("Number 1 is not a prime")
		return
	}

	isPrime := true
	sqrt := int(math.Sqrt(float64(number)))
	for i := 2; i < sqrt; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}

	fmt.Printf("The number %d is prime? %t", number, isPrime)
}
