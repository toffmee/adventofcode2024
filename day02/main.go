package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	filename := "input.txt"

	start := time.Now()
	part1Result := part1(filename)
	fmt.Printf("Part 1 answer: %d in %v\n", part1Result, time.Since(start))

	start = time.Now()
	part2Result := part2(filename)
	fmt.Printf("Part 2 answer: %d in %v\n", part2Result, time.Since(start))
}

func part1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var safe int
	var data [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		levels := make([]int, len(fields))

		for i, str := range fields {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			levels[i] = num
		}
		data = append(data, levels)
	}

	for _, levels := range data {
		if checkSafety(levels) {
			safe++
		}
	}

	return safe
}

func checkSafety(levels []int) bool {
	var increasing, decreasing bool

	diff := levels[1] - levels[0]
	if diff > 0 {
		increasing = true
	} else if diff < 0 {
		decreasing = true
	} else {
		return false
	}

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		absDiff := diff
		if diff < 0 {
			absDiff = -diff
		}

		if absDiff < 1 || absDiff > 3 {
			return false
		}

		if increasing && diff <= 0 {
			return false
		}

		if decreasing && diff >= 0 {
			return false
		}
	}

	return true

}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var safe int
	var data [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		levels := make([]int, len(fields))

		for i, str := range fields {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			levels[i] = num
		}
		data = append(data, levels)
	}

	for _, levels := range data {
		if checkSafetyWithRemoval(levels) {
			safe++
		}
	}

	return safe
}

func checkSafetyWithRemoval(levels []int) bool {
	if checkSafety(levels) {
		return true
	}

	for i := 0; i < len(levels); i++ {
		newLevels := make([]int, 0, len(levels)-1)
		newLevels = append(newLevels, levels[:i]...)
		newLevels = append(newLevels, levels[i+1:]...)

		if checkSafety(newLevels) {
			return true
		}
	}

	return false
}
