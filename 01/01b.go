package aoc01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// func main() {
// 	result := countRangeIncreases("sample", 3)

// 	fmt.Println(result)
// }

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
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("failed to parse value: %s", err)
		}

		fmt.Printf("\nrow %d value %d\n", row, value)

		completeIndex := row % rangeSize
		compareIndex := (row + 1) % rangeSize
		complete := measurements[completeIndex]
		measurements[completeIndex] = make([]int, 0)
		fmt.Printf("complete %v\n", complete)

		for i, _ := range measurements {
			measurements[i] = append(measurements[i], value)
			fmt.Printf("%d %v\n", i, measurements[i])
		}

		if len(complete) == rangeSize && len(measurements[compareIndex]) == rangeSize {
			completeAvg := sum(complete)
			compareAvg := sum(measurements[compareIndex])

			if compareAvg > completeAvg {
				result = result + 1
			}
		}

		row = row + 1
	}

	file.Close()

	return result
}

func sum(s []int) int {
	var total = 0
	for _, v := range s {
		total = total + v
	}
	return total
}
