package main

import "fmt"

func main() {
	var elements int
	fmt.Print("Enter the total number of elements: ")
	fmt.Scanln(&elements)

	slice := unSorted(elements)
	fmt.Println("\n--- Unsorted element --- \n", slice)
	bubblesort(slice)
	fmt.Println("\n--- Sorted element ---\n", slice, "\n")
}

func unSorted(size int) []int {
	var element int

	slice := make([]int, size, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter the element %d: ", i+1)
		fmt.Scanln(&element)
		slice[i] = element
	}

	return slice
}

func bubblesort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
