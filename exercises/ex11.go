package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*Write a Go program that takes a slice of integers as input and returns a new slice containing only the unique elements from the original slice, in the order they appeared.

Use the fmt package to prompt the user to enter a series of integers, separated by spaces.

Ex:
Input: 3 5 2 5 7 3 8
Unique elements: [3 5 2 7 8]

You can solve it using a map or only using slices.
*/

func main() {
	inputArr, err := readInputArray()
	if err != nil {
		fmt.Println("Error when reading array", err)
		return
	}

	uniqueElements := make([]int, 0)
	seenElements := make(map[int]bool)
	for _, v := range inputArr {
		if _, ok := seenElements[v]; !ok {
			uniqueElements = append(uniqueElements, v)
			seenElements[v] = true
		}
	}

	fmt.Println(uniqueElements)
}

func readInputArray() ([]int, error) {
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
			return nil, errors.New("user should inform only integers")
		}

		numbers[i] = converted
	}

	return numbers, nil
}
