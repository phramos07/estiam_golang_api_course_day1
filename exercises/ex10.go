package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a Go program that takes a series of numbers separated by spaces as input and calculates the average of those numbers using slices.

Ex:
Input: 10 20 30 40 50
Output: 30

Tip: To read a series of space-separated numbers, you can use the fmt.Scanln function and then use the strings.Fields function to split the input into individual numbers.

// Tips:
- Convert each string to int
	- If conversion not possible, handle the error
- Calculate the average
*/

//strconv.Atoi -> use this function to convert

func main() {
	var input string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input the array:")
	if scanner.Scan() {
		input = scanner.Text()
	}

	separated := strings.Fields(input)
	numbers := make([]int, len(separated))
	for i, v := range separated {
		converted, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println("User should inform only integers.", err)
			return
		}

		// numbers = append(numbers, converted)
		numbers[i] = converted
	}

	sum := 0
	for _, v := range numbers {
		sum += v
	}

	fmt.Printf("%.2f", float64(sum)/float64(len(numbers)))
}
