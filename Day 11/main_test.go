package main

import (
	"testing"
)

func TestPart1a(t *testing.T) {
	total := part1([]string{"11111", "19991", "19191", "19991", "11111"}, 2)
	if total != 9 {
		t.Errorf("Part 1a was incorrect, got %d, want %d,", total, 9)
	}
}

/*
func TestPart1b(t *testing.T) {
	total := part1([]string{"5483143223", "2745854711", "5264556173", "6141336146", "6357385478", "4167524645", "2176841721", "6882881134", "4846848554", "5283751526"}, 10)
	if total != 204 {
		t.Errorf("Part 1b was incorrect, got %d, want %d,", total, 204)
	}
}

func TestPart1c(t *testing.T) {
	total := part1([]string{"5483143223", "2745854711", "5264556173", "6141336146", "6357385478", "4167524645", "2176841721", "6882881134", "4846848554", "5283751526"}, 100)
	if total != 1656 {
		t.Errorf("Part 1c was incorrect, got %d, want %d,", total, 1656)
	}
}

(func TestPart2(t *testing.T) {
	total := part2([]string{"[({(<(())[]>[[{[]{<()<>>","[(()[<>])]({[<{<<[]>>(","{([(<{}[<>[]}>{[]{[(<()>","(((({<>}<{<{<>}{[]{[]{}","[[<[([]))<([[{}[[()]]]","[{[{({}]{}}([{[{{{}}([]","{<[[]]>}<{[{[{[]{()[[[]","[<(<(<(<{}))><([]([]()","<{([([[(<>()){}]>(<<{{","<{([{{}}[<[[[<>{}]]]>[]]"})
	if total != 288957 {
		t.Errorf("Part 2 was incorrect, got %d, want %d,", total, 288957)
	}
}


*/
