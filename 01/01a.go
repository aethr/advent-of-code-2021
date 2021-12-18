package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	scanner.Scan()
	last, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("failed to parse value: %s", err)
	}

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("failed to parse value: %s", err)
		}
		if value > last {
			result = result + 1
		}
		last = value
	}

	file.Close()

	fmt.Println(result)
}
