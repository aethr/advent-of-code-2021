package aoc02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getPositionAndDepthUsingAim(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	position := 0
	depth := 0
	aim := 0

	for scanner.Scan() {
		position, depth, aim = adjustPositionAndDepthAndAim(scanner.Text(), position, depth, aim)
	}

	file.Close()
	return position, depth
}

func adjustPositionAndDepthAndAim(command string, position int, depth int, aim int) (int, int, int) {
	parts := strings.Fields(command)
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("failed to convert value to int: %s", err)
	}

	switch parts[0] {
	case "forward":
		position = position + value
		depth = depth + (aim * value)
	case "up":
		aim = aim - value

	case "down":
		aim = aim + value
	}
	return position, depth, aim
}
