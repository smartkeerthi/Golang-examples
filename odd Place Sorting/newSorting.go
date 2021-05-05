package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{9, 7, 1, 2, 0, 5}
	fmt.Print(input)

	newSorting(input)
}

func newSorting(input []int) {
	length := len(input)
	inputArr := make([]int, length, length)
	var oddNum []int
	var evenNum []int

	for i := 0; i < length; i++ {
		if i%2 == 0 {
			evenNum = append(evenNum, input[i])
		} else {
			oddNum = append(oddNum, input[i])
		}
	}
	sort.Ints(evenNum)
	sort.Sort(sort.Reverse(sort.IntSlice(oddNum)))

	evenCount := 0
	oddCount := 0
	for j := 0; j < len(inputArr); j++ {
		if j%2 == 0 {
			inputArr[j] = evenNum[evenCount]
			evenCount = evenCount + 1
		} else {
			inputArr[j] = oddNum[oddCount]
			oddCount = oddCount + 1
		}
	}

	fmt.Print(inputArr)
}
