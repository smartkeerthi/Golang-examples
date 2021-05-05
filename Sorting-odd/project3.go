// Sorting only the odd numbers in the array

package main

import (
	"fmt"
	"sort"
)

func main() {
	var input = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("Input: ", input)
	newSorting(input)
}

func newSorting(input []int) {
	var output []int
	output = input
	var inputLen = len(output)
	var oddNum []int

	for i := 0; i < inputLen; i++ {
		if input[i]%2 != 0 {
			oddNum = append(oddNum, input[i])
			input[i] = 1
		}
	}
	sort.Ints(oddNum)
	var c = 0
	for i := 0; i < len(output); i++ {
		if output[i] == 1 {
			output[i] = oddNum[c]
			c = c + 1
		}
	}
	fmt.Println("output:", output)
}
