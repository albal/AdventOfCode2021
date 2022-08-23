package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	x, y int
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("Day 10\\input.txt")
	if err != nil {
		fmt.Println("Error parsing value")
	}
	fmt.Printf("Part1: %d\n", part1(lines))
	fmt.Printf("Part2: %d\n", part2(lines))
}

func part1(values []string) int {
	var errors []int32
	for _, line := range values {
		var stack Stack
		for _, char := range line {
			if isOpen(char) {
				stack.Push(char)
			}
			if isClose(char) {
				val, _ := stack.Pop()
				if isNotOpposite(val, char) {
					errors = append(errors, char)
					break
				}
			}
		}
	}
	return calculateScore(errors)
}

func part2(values []string) int64 {
	var errors []int32
	var incomplete []string
	var scores []int64
	for _, line := range values {
		var stack Stack
		var added string
		var corrupted bool = true
		for _, char := range line {
			if isOpen(char) {
				stack.Push(char)
			}
			if isClose(char) {
				val, _ := stack.Pop()
				if isNotOpposite(val, char) {
					errors = append(errors, char)
					corrupted = false
					break
				}
			}
		}
		if corrupted {
			for i := 0; i < len(stack); i++ {
				c := stack[len(stack)-1-i]
				line = line + string(opposite(c))
				added = added + string(opposite(c))
			}
			incomplete = append(incomplete, line)
			scores = append(scores, calcSymbolScore(added))
		}
	}

	for _, val := range scores {
		fmt.Printf("%d\n", val)
	}

	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })

	return scores[(len(scores) / 2)]
}

func calcSymbolScore(vals string) int64 {
	var score int64 = 0
	for _, val := range vals {
		score *= 5
		value := int64(val)
		score += lookup(value)
	}
	return score
}

func lookup(val int64) int64 {
	var res int64
	switch val {
	case ')':
		res = 1
	case ']':
		res = 2
	case '}':
		res = 3
	case '>':
		res = 4
	default:
		res = 0
		fmt.Printf("Error we had %c", val)
	}
	return res
}

func opposite(val int32) int32 {
	var res int32
	switch val {
	case '(':
		res = ')'
	case '[':
		res = ']'
	case '{':
		res = '}'
	case '<':
		res = '>'
	default:
		res = ' '
		fmt.Printf("Error we had %c", val)
	}
	return res
}
func isNotOpposite(val int32, char int32) bool {
	if val == '(' && char == ')' {
		return false
	}
	if val == '{' && char == '}' {
		return false
	}
	if val == '[' && char == ']' {
		return false
	}
	if val == '<' && char == '>' {
		return false
	}
	return true
}

func isOpen(char int32) bool {
	if char == '{' || char == '(' || char == '[' || char == '<' {
		return true
	}
	return false
}

func isClose(char int32) bool {
	if char == '}' || char == ')' || char == ']' || char == '>' {
		return true
	}
	return false
}

func calculateScore(errors []int32) int {
	/*
		): 3 points.
		]: 57 points.
		}: 1197 points.
		>: 25137 points.
	*/
	count := 0
	for _, err := range errors {
		switch err {
		case ')':
			count += 3
		case ']':
			count += 57
		case '}':
			count += 1197
		case '>':
			count += 25137
		default:
			fmt.Printf("Got %c but can't process it", err)
		}
	}
	return count
}

type Stack []int32

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(char int32) {
	*s = append(*s, char) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int32, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
