package main

import "fmt"

/*Write a Go program that converts a temperature from Celsius to Fahrenheit using a function.

fahrenheit = (celsius * 9/5) + 32
*/

func main() {
	c := 40.0
	fmt.Printf("From celsius %f to fahrenheit %f", c, fromCelsiusToFahrenheit(c))
}

func fromCelsiusToFahrenheit(celsius float64) (fahrenheit float64) { //named return
	fahrenheit = (celsius * 9.0 / 5.0) + 32.0
	return
}
