package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	total := part1([]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"})
	if total != 198 {
		t.Errorf("Part 1 was incorrect, got %d, want %d,", total, 198)
	}
}

func TestPart2(t *testing.T) {
	total := part2([]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"})
	if total != 230 {
		t.Errorf("Part 1 was incorrect, got %d, want %d,", total, 230)
	}
}
