package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers, err := readLines("assets/input")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	fmt.Printf(
		"Part 1: We found %d out of %d numbers larger than their respective predecessor\n",
		countLargerThanPredecessor(numbers),
		len(numbers),
	)

	fmt.Printf(
		"Part 2: We found %d out of %d measurement-windows larger than their respective predecessor-windows\n",
		countLargerThanPrecedingWindowOfTwo(numbers),
		len(numbers)-2,
	)
}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return lines, err
		}
		lines = append(lines, number)
	}
	return lines, scanner.Err()
}

func countLargerThanPredecessor(numbers []int) int {
	var predecessor int
	result := 0

	for i, currentNumber := range numbers {
		if i > 0 && predecessor < currentNumber {
			result++
		}
		predecessor = currentNumber
	}

	return result
}

func countLargerThanPrecedingWindowOfTwo(numbers []int) int {
	var windowSums []int
	for i, currentNumber := range numbers {
		if i < 2 {
			continue
		}

		windowSums = append(windowSums, currentNumber+numbers[i-1]+numbers[i-2])
	}

	return countLargerThanPredecessor(windowSums)
}
