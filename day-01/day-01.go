package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr1, arr2 := getArrays("day-01.txt")

	// ### Q2
	occurence := make(map[int]int)
	for i := 0; i < len(arr1); i++ {
		occurence[arr1[i]] += +1
	}
	score := 0
	for i := 0; i < len(arr2); i++ {
		score += occurence[arr2[i]] * arr2[i]
	}
	fmt.Println(score)

	// ### Q1
	// sort.Ints(arr1)
	// sort.Ints(arr2)
	// diff := 0
	// for i := 0; i < len(arr1); i++ {
	// 	diff += abs(arr1[i] - arr2[i])
	// }
	// fmt.Println(diff)
}

func getArrays(path string) ([]int, []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Fields(line)
		// parse and append values value
		arr1 = parseAndAppendValue(splitLine[0], arr1)
		arr2 = parseAndAppendValue(splitLine[1], arr2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return arr1, arr2
}

func parseAndAppendValue(str string, arr []int) []int {
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return append(arr, val)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
