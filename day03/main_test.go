package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 161
	result := part1("test1.txt")
	if result != expected {
		t.Errorf("part1() = %d; want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 48
	result := part2("test2.txt")
	if result != expected {
		t.Errorf("part2() = %d; want %d", result, expected)
	}
}
