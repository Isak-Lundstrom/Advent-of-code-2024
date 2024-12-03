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
	path := "day-02.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	num := 0

	for scanner.Scan() {
		report := scanner.Text()
		levels := strings.Fields(report)
		fmt.Println(levels)
		for i := 0; i < len(levels); i++ {
			levelMinusBad := make([]string, 0, len(levels)-1)
			levelMinusBad = append(levelMinusBad, levels[:i]...)
			levelMinusBad = append(levelMinusBad, levels[i+1:]...)
			if isReportSafe(levelMinusBad) {
				num++
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)
}

func isReportSafe(levels []string) bool {
	var inc bool
	var prev int
	for i, l := range levels {
		level, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		// if first only set prev
		if i == 0 {
			prev = level
			continue
		}
		// diff has to be at least 1 but not more than 3
		diff := abs(level - prev)
		if diff < 1 || diff > 3 {
			return false
		}
		// check increasing or decreasing
		if i == 1 {
			inc = level < prev
		}
		// continue if first 2 levels
		if i < 2 {
			prev = level
			continue
		}
		// all levels must increase or decrease
		if inc == (level > prev) {
			return false
		}
		prev = level
	}
	return true
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// func main() {
// 	path := "day-02.txt"
// 	file, err := os.Open(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	num := 0

// 	for scanner.Scan() {
// 		var inc bool
// 		var prev int
// 		report := scanner.Text()
// 		levels := strings.Fields(report)
// 		for i, l := range levels {
// 			level, err := strconv.Atoi(l)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			if i == 0 {
// 				prev = level
// 				continue
// 			}
// 			// diff has to be at least 1 but not more than 3
// 			diff := abs(level - prev)
// 			if diff < 1 || diff > 3 {
// 				num--
// 				break
// 			}
// 			// check increasing or decreasing
// 			if i == 1 {
// 				inc = level < prev
// 			}
// 			// continue if first 2 levels
// 			if i < 2 {
// 				prev = level
// 				continue
// 			}
// 			// all levels must increase or decrease
// 			if inc == (level > prev) {
// 				num--
// 				break
// 			}
// 			prev = level
// 		}
// 		num++
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(num)
// }

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
