package main

import (
	"bufio"
	"fmt"
	"os"
)

/*Write a Go program that reads a file and counts the number of lines, handling file reading errors.


Example:
Input: input.txt
Number of lines in the file: 10

Input: non_existent.txt
Error: open non_existent.txt: no such file or directory

If the file is not there or any error happens this should be handled.

Tips: Implement the countLines function to handle file opening, line counting, and error handling using os.Open and bufio.Scanner.
*/

func main() {
	filename, err := readFilename()
	if err != nil {
		fmt.Println("Error when reading the filename from input", err)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error when opening the file", err)
		return
	}

	linecount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linecount++
	}

	fmt.Printf("Filename: %s\n", filename)
	fmt.Printf("Number of lines: %d\n", linecount)

}

func readFilename() (string, error) {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input the filename:")
	if scanner.Scan() {
		input = scanner.Text()
	} else {
		return "", fmt.Errorf("Error when scanning. Input: %s", input)
	}

	return input, nil
}
