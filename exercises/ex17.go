package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/shopspring/decimal"
)

// To installl decimal lib, start a go module

// go mod init [nameofthemodule]
// add decimal as dependency on go.mod file OR in any .go file
// run go mod tidy (unused dependencies will be removed! uninstalled dependencies will be installed.)

/*Write a Go program that performs calculations on decimal numbers with high precision. Use the shopstring library

Install the lib:

go get github.com/shopspring/decimal

Ex:
Enter the first decimal number: 123.456
Enter the second decimal number: 789.012
Select operation (add, subtract, multiply, divide): divide
Result: 0.15624672430556339

Tips:
convert from string to decimal
Selecting operation - switch
Print result with maximum precision possible
*/

func main() {
	dec1, err := readDecimal()
	if err != nil {
		fmt.Printf("Error when reading decimal number: %s", err.Error())
		return
	}

	// get the second number (reuse code)
	dec2, err := readDecimal()
	if err != nil {
		fmt.Printf("Error when reading decimal number: %s", err.Error())
		return
	}

	// select operation
	op, err := readOperation()
	if err != nil {
		fmt.Printf("Error when reading operation: %s", err.Error())
		return
	}

	// switch on operation
	var result decimal.Decimal
	switch op {
	case "ADD":
		result = dec1.Add(*dec2)
	case "DIV":
		result = dec1.Div(*dec2)
	case "MUL":
		result = dec1.Mul(*dec2)
	case "SUB":
		result = dec1.Sub(*dec2)
	}

	// print result
	fmt.Println(result)
}

func readDecimal() (*decimal.Decimal, error) {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input a number: ")
	if scanner.Scan() {
		input = scanner.Text()
	}

	dec, err := decimal.NewFromString(input)
	if err != nil {
		fmt.Printf("Error when reading decimal number: %s", err.Error())
		return nil, err
	}

	return &dec, nil
}

func readOperation() (string, error) {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Select the operation:\nADD\nDIV\nMUL\nSUB\n")
	if scanner.Scan() {
		input = scanner.Text()
	} else {
		return "", fmt.Errorf("Error when scanning. Input: %s", input)
	}

	return input, nil
}
