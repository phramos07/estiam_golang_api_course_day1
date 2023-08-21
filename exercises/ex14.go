package main

import (
	"fmt"
)

/*Write a Go program that reverses the elements of an array.

Original array: [1 2 3 4 5]
Reversed array: [5 4 3 2 1]

You can use the package sort.
*/

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	reversedCopy, err := reverseArrayCopy(arr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Copy: ", reversedCopy)

	reverseArrayInplace(&arr)
	fmt.Println("In place: ", arr)
}

// Questions:
// make a copy of the array?
// reverse inplace (don't make a copy)?
// return array or just alter the argument?
// receive a reference or a copy as argument?
// func reversesArray(???) ???

// sorting in place without allocating new memory
// no copy - sort inplace within the function
func reverseArrayInplace(input *[5]int) { // I am copying the whole array
	// reverse without sort and iterating over the array ONLY ONCE.
	arrLen := len(*input)
	for i := 0; i < arrLen/2; i++ {
		(*input)[i], (*input)[arrLen-1-i] = (*input)[arrLen-1-i], (*input)[i]
	}
} // O(n) -> optimal swap

// * -> access the contents of an address OR defines a pointer type
// & -> access the address of a variable

// copy - return new slice
func reverseArrayCopy(arr [5]int) ([]int, error) {
	revArray := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		revArray[i] = arr[len(arr)-i-1]
	}
	return revArray, nil
}
