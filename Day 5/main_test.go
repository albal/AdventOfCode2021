package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	total := part1([]string{"0,9 -> 5,9", "8,0 -> 0,8", "9,4 -> 3,4", "2,2 -> 2,1", "7,0 -> 7,4", "6,4 -> 2,0", "0,9 -> 2,9", "3,4 -> 1,4", "0,0 -> 8,8", "5,5 -> 8,2"})
	if total != 5 {
		t.Errorf("Part 1 was incorrect, got %d, want %d,", total, 5)
	}
}

func TestPart2(t *testing.T) {
	total := part2([]string{"0,9 -> 5,9", "8,0 -> 0,8", "9,4 -> 3,4", "2,2 -> 2,1", "7,0 -> 7,4", "6,4 -> 2,0", "0,9 -> 2,9", "3,4 -> 1,4", "0,0 -> 8,8", "5,5 -> 8,2"})
	if total != 12 {
		t.Errorf("Part 2 was incorrect, got %d, want %d,", total, 12)
	}
}
