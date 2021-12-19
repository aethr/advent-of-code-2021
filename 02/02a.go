package aoc02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getPositionAndDepth(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	position := 0
	depth := 0

	for scanner.Scan() {
		position, depth = adjustPositionAndDepth(scanner.Text(), position, depth)
	}

	file.Close()
	return position, depth
}

func adjustPositionAndDepth(command string, position int, depth int) (int, int) {
	parts := strings.Fields(command)
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("failed to convert value to int: %s", err)
	}
	switch parts[0] {
	case "forward":
		position = position + value

	case "up":
		depth = depth - value

	case "down":
		depth = depth + value
	}
	return position, depth
}
