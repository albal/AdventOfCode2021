package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	total := part1([]int64{16, 1, 2, 0, 4, 2, 7, 1, 2, 14})
	if total != 37 {
		t.Errorf("Part 1 was incorrect, got %d, want %d,", total, 37)
	}
}

func TestPart2(t *testing.T) {
	total := part2([]int64{16, 1, 2, 0, 4, 2, 7, 1, 2, 14})
	if total != 168 {
		t.Errorf("Part 2 was incorrect, got %d, want %d,", total, 168)
	}
}

func TestMove1(t *testing.T) {
	total := move(16, 5)
	if total != 66 {
		t.Errorf("move() Test 1 was incorrect, got %d, want %d,", total, 66)
	}
}

func TestMove2(t *testing.T) {
	total := move(14, 5)
	if total != 45 {
		t.Errorf("move() Test 2 was incorrect, got %d, want %d,", total, 45)
	}
}

func TestMove3(t *testing.T) {
	total := move(1, 5)
	if total != 10 {
		t.Errorf("move() Test 3 was incorrect, got %d, want %d,", total, 10)
	}
}
