package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	path := "day-04.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var matrix []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		matrix = append(matrix, row)
	}

	total1 := q1(matrix)
	total2 := q2(matrix)
	fmt.Println(total1, " ", total2)
}

func q2(matrix []string) int {
	total := 0
	rowLen := len(matrix)
	colLen := len(matrix[0])
	for rowIndex := 1; rowIndex < rowLen-1; rowIndex++ {
		for colIndex := 1; colIndex < colLen-1; colIndex++ {
			if matrix[rowIndex][colIndex] != 'A' {
				continue
			}
			if checkX_MAS(matrix, rowIndex, colIndex, 'M', 'M', 'S', 'S') ||
				checkX_MAS(matrix, rowIndex, colIndex, 'S', 'M', 'M', 'S') ||
				checkX_MAS(matrix, rowIndex, colIndex, 'S', 'S', 'M', 'M') ||
				checkX_MAS(matrix, rowIndex, colIndex, 'M', 'S', 'S', 'M') {
				total += 1
			}
		}
	}
	return total
}

func checkX_MAS(matrix []string, r, c int, c1, c2, c3, c4 rune) bool {
	return rune(matrix[r+1][c-1]) == c1 && rune(matrix[r-1][c-1]) == c2 && rune(matrix[r-1][c+1]) == c3 && rune(matrix[r+1][c+1]) == c4
}

func q1(matrix []string) int {
	total := 0
	rowLen := len(matrix)
	colLen := len(matrix[0])
	for rowIndex := 0; rowIndex < rowLen; rowIndex++ {
		upOK := rowIndex-3 >= 0
		downOK := rowIndex+3 < rowLen
		for colIndex := 0; colIndex < colLen; colIndex++ {
			if matrix[rowIndex][colIndex] != 'X' {
				continue
			}
			rightOK := colIndex+3 < colLen
			leftOK := colIndex-3 >= 0

			if downOK && checkMAS(matrix, rowIndex+1, colIndex, rowIndex+2, colIndex, rowIndex+3, colIndex) {
				total += 1
			}

			if downOK && rightOK && checkMAS(matrix, rowIndex+1, colIndex+1, rowIndex+2, colIndex+2, rowIndex+3, colIndex+3) {
				total += 1
			}

			if rightOK && checkMAS(matrix, rowIndex, colIndex+1, rowIndex, colIndex+2, rowIndex, colIndex+3) {
				total += 1
			}

			if upOK && rightOK && checkMAS(matrix, rowIndex-1, colIndex+1, rowIndex-2, colIndex+2, rowIndex-3, colIndex+3) {
				total += 1
			}

			if upOK && checkMAS(matrix, rowIndex-1, colIndex, rowIndex-2, colIndex, rowIndex-3, colIndex) {
				total += 1
			}

			if upOK && leftOK && checkMAS(matrix, rowIndex-1, colIndex-1, rowIndex-2, colIndex-2, rowIndex-3, colIndex-3) {
				total += 1
			}

			if leftOK && checkMAS(matrix, rowIndex, colIndex-1, rowIndex, colIndex-2, rowIndex, colIndex-3) {
				total += 1
			}

			if downOK && leftOK && checkMAS(matrix, rowIndex+1, colIndex-1, rowIndex+2, colIndex-2, rowIndex+3, colIndex-3) {
				total += 1
			}

		}
	}
	return total
}

func checkMAS(matrix []string, r1, c1, r2, c2, r3, c3 int) bool {
	return matrix[r1][c1] == 'M' && matrix[r2][c2] == 'A' && matrix[r3][c3] == 'S'
}
