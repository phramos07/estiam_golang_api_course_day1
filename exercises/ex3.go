package main

import "fmt"

/*Strings

Take the sentence "Golang is an awesome language" and iterate over each character, printing the index, the character and the ascii value (bytes).
*/

func main() {
	sentence := "Golang is an awesome language"

	// implementing theo's approach - traditional for

	// Golang only has the for loop
	for i := 0; i < len(sentence); i++ {
		fmt.Printf("Index: %d, value: %c, ASCII: %d\n", i, sentence[i], sentence[i])
	}

	// using range
	for i, v := range sentence {
		fmt.Printf("Index: %d, value: %c, ASCII: %d\n", i, v, v)
	}

	// implementing Ruben's approach (while str[i] != '\0')
	// i := 0
	// for sentence[i] != "\0" {
	// 	fmt.Printf("Index: %d, value: %c, ASCII: %d\n", i, sentence[i], sentence[i])
	// 	i++
	// }
}
