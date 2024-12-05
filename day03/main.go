package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	content := string(data)

	re := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	matches := re.FindAllStringSubmatch(content, -1)

	sum := 0
	for _, m := range matches {
		x, errX := strconv.Atoi(m[1])
		y, errY := strconv.Atoi(m[2])
		if errX == nil && errY == nil {
			sum += x * y
		}
	}

	return sum
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	content := string(data)

	re := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)|do\\(\\)|don't\\(\\)")

	matches := re.FindAllStringSubmatch(content, -1)

	mulEnabled := true
	sum := 0

	for _, m := range matches {
		fullMatch := m[0]

		if strings.HasPrefix(fullMatch, "mul(") {
			if mulEnabled {
				x, errX := strconv.Atoi(m[1])
				y, errY := strconv.Atoi(m[2])
				if errX == nil && errY == nil {
					sum += x * y
				}
			}
		} else if fullMatch == "do()" {
			mulEnabled = true
		} else if fullMatch == "don't()" {
			mulEnabled = false
		}
	}

	return sum
}
