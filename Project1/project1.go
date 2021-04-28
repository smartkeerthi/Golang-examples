package main

import "fmt"

func main() {
	var matrixA [100][100]int
	var matrixB [100][100]int
	var sum [100][100]int
	var rows, cols int

	fmt.Print("Number of rows: ")
	fmt.Scanln(&rows)
	fmt.Print("Number of Columns: ")
	fmt.Scanln(&cols)

	fmt.Println("========== Matrix A ==========")

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("Enter the value for matrixA %d %d: ", i+1, j+1)
			fmt.Scanln(&matrixA[i][j])
		}
	}

	fmt.Println("========== Matrix B ==========")

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("Enter the value for matrixB %d %d: ", i+1, j+1)
			fmt.Scanln(&matrixB[i][j])
		}
	}

	fmt.Println("========== SUM ==========")

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			sum[i][j] = matrixA[i][j] + matrixB[i][j]
			fmt.Printf(" %d ", sum[i][j])
			if j == cols-1 {
				fmt.Println("")
			}
		}
	}
}
