package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	total := part1([]string{"2199943210", "3987894921", "9856789892", "8767896789", "9899965678"})
	if total != 15 {
		t.Errorf("Part 1 was incorrect, got %d, want %d,", total, 15)
	}
}

func TestPart2(t *testing.T) {
	total := part2([]string{"2199943210", "3987894921", "9856789892", "8767896789", "9899965678"})
	if total != 1134 {
		t.Errorf("Part 2 was incorrect, got %d, want %d,", total, 1134)
	}
}
