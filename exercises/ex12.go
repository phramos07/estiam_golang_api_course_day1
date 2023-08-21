package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a Go program that takes a string as input and counts the frequency of each word in the input.

1. Use the fmt package to prompt the user to enter a string.
2. Implement a program that processes the input string, splits it into words, and counts the frequency of each word using a map.

ex:

Input: the quick brown fox jumps over the lazy dog the lazy dog

Word frequency count:
the: 3
quick: 1
brown: 1
fox: 1
jumps: 1
over: 1
lazy: 2
dog: 2

Tip: You can use the strings.Fields function to split the input string into words.
*/

func main() {
	inputArr, err := readInputArray()
	if err != nil {
		fmt.Println("Error when reading array", err)
		return
	}

	wordCount := make(map[string]int, 0)
	for _, v := range inputArr {
		wordCount[v]++ // in the case of a map[x]int, int zero value is 0, so this works without checking if the key exists on the map
	}

	for key, value := range wordCount {
		fmt.Printf("word: %s, count: %d\n", key, value)
	}
}

func readInputArray() ([]string, error) {
	var input string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input the array:")
	if scanner.Scan() {
		input = scanner.Text()
	} else {
		return nil, fmt.Errorf("Error when scanning. Input: %s", input) // errorF allows you to format the output. in this case input is always empty (reaching definition)
	}

	return strings.Fields(input), nil
}
