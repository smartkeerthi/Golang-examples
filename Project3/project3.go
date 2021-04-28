package main

import "fmt"

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
	odd := make([]int, itemsLen, itemsLen)
	var c = 0

	for i := 0; i < itemsLen; i++ {
		if items[i]%2 == 0 {
			fmt.Println("Even")
		} else {
			odd[c] = items[i]
			c = c + 1
			items[i] = 0
		}
	}
	fmt.Printf("item, %d", items)
	fmt.Printf("odd, %d", odd)
}
