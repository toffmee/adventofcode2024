package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	var first, second []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		firstItem, _ := strconv.Atoi(parts[0])
		secondItem, _ := strconv.Atoi(parts[1])
		first = append(first, firstItem)
		second = append(second, secondItem)
	}

	sort.Ints(first)
	sort.Ints(second)

	sumOfDifferences := 0
	for i := 0; i < len(first) && i < len(second); i++ {
		diff := first[i] - second[i]
		if diff < 0 {
			diff = -diff
		}
		sumOfDifferences += diff
	}

	return sumOfDifferences
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var first, second []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		firstItem, _ := strconv.Atoi(parts[0])
		secondItem, _ := strconv.Atoi(parts[1])
		first = append(first, firstItem)
		second = append(second, secondItem)
	}

	countMap := make(map[int]int)
	for _, num := range second {
		countMap[num]++
	}

	similarityScore := 0
	for _, num := range first {
		similarityScore += num * countMap[num]
	}

	return similarityScore
}
