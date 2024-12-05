package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	path := "day-05.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Section 1
	entries := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			// next section
			break
		}
		pair := strings.Split(line, `|`)
		entries[pair[1]] = append(entries[pair[1]], pair[0])
	}
	fmt.Println(entries)

	// Section 2
	var instructions [][]string
	for scanner.Scan() {
		lineOk := true
		line := scanner.Text()
		values := strings.Split(line, `,`)
		for i, value := range values {
			if i == 0 {
				continue
			}
			entriesList := entries[value]
			for _, v := range values[:i] {
				if slices.Contains(entriesList, v) {
					continue
				}
				lineOk = false
			}
		}

		if lineOk {
			instructions = append(instructions, values)
		}
	}
	fmt.Println(instructions)

	total := 0
	for _, ins := range instructions {
		val, err := strconv.Atoi(ins[len(ins)/2])
		if err != nil {
			log.Fatal(err)
		}
		total += val
	}
	fmt.Println(total)

}
