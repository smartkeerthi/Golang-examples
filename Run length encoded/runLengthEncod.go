package main

import (
	"fmt"
)

func main() {
	input := "aaaabbbccdddddeeeeeeeee"
	runLength(input)
}

func runLength(input string) {
	length := len(input)
	for i := 0; i < length; i++ {
		count := 1
		for i < length-1 && input[i] == input[i+1] {
			count = count + 1
			i = i + 1
		}
		fmt.Print(string(input[i]), count)
	}
}
