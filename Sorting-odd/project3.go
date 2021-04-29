// Sorting only the odd numbers in the array

package main

import (
	"fmt"
	"sort"
)

func main() {
	var totalElement int

	fmt.Print("Enter the total element : ")
	fmt.Scanln(&totalElement)

	elements := unSorted(totalElement)
	fmt.Println("elements ", elements)
	newSorting(elements)
}

func unSorted(size int) []int {
	var element int

	elements := make([]int, size, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter the element %d: ", i+1)
		fmt.Scanln(&element)
		elements[i] = element
	}

	return elements
}

func newSorting(items []int) {
	var temp []int
	temp = items
	var itemsLen = len(temp)
	var odd []int

	for i := 0; i < itemsLen; i++ {
		if items[i]%2 == 0 {
			fmt.Println("Even")
		} else {
			odd = append(odd, items[i])
			items[i] = 1
		}
	}
	sort.Ints(odd)
	sorting(temp, odd)
}
func sorting(temp, odd []int) {
	var c = 0
	for i := 0; i < len(temp); i++ {
		if temp[i] == 1 {
			temp[i] = odd[c]
			c = c + 1
		}
	}
	fmt.Println("temp:", temp)
}
