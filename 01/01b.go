package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	result := countRangeIncreases("sample", 3)

	fmt.Println(result)
}

func countRangeIncreases(filename string, rangeSize int) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	row := 0
	measurements := make([][]int, rangeSize)

	if err != nil {
		log.Fatalf("failed to parse value: %s", err)
	}

	for scanner.Scan() {
		window := row % rangeSize

		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("failed to parse value: %s", err)
		}

		fmt.Printf("\nrow %d value %d\n", row, value)

		for i, _ := range measurements {
			if i >= row {
				measurements[(i+window)%rangeSize] = append(measurements[i], value)
			}
			fmt.Printf("%d %s\n", i, measurements[i])
		}

		// if value > last {
		// 	result = result + 1
		// }

		row = row + 1
	}

	file.Close()

	return result
}
