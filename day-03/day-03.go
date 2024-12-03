package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	path := "day-03.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	do := true

	for scanner.Scan() {
		sum := 0
		index := 0
		dIndex := -1
		mIndex := -1
		pIndex := -1
		row := scanner.Text()
		for index < len(row) {
			// Pointers for start and end of commands
			if row[index] == 'd' {
				dIndex = index
			}
			if row[index] == 'm' {
				mIndex = index
			}
			if row[index] == ')' {
				pIndex = index
			}
			// Check mul command
			if pIndex > mIndex && mIndex > 0 && pIndex > 0 && index <= pIndex {
				reMul := regexp.MustCompile(`^mul\(([0-9]+),([0-9]+)\)$`)
				matchMul := reMul.FindStringSubmatch(row[mIndex : pIndex+1])

				if len(matchMul) == 3 {
					num1, err1 := strconv.Atoi(matchMul[1])
					num2, err2 := strconv.Atoi(matchMul[2])
					if do && err1 == nil && err2 == nil {
						sum += num1 * num2
					}
				}
			}
			// Check do and don't command
			if pIndex > dIndex && dIndex > 0 && pIndex > 0 && index <= pIndex {
				reDo := regexp.MustCompile(`(?i)do\(\)`)
				reDont := regexp.MustCompile(`(?i)don't\(\)`)
				matchDo := reDo.MatchString(row[dIndex : pIndex+1])
				matchDont := reDont.MatchString(row[dIndex : pIndex+1])

				if matchDo {
					do = true
				}

				if matchDont {
					do = false
				}
			}
			index++
		}
		total += sum
	}
	fmt.Println(total)
}
