package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	lines, err := readLines("Day 11\\input.txt")
	if err != nil {
		fmt.Println("Error parsing value")
	}
	fmt.Printf("Part1: %d\n", part1(lines, 100))
	//fmt.Printf("Part2: %d\n", part2(lines))
}

func part1(values []string, turns int) int {
	var flashes int = 0
	var octo [][]int = make([][]int, len(values))
	for i, line := range values {
		for _, val := range line {
			num, _ := strconv.Atoi(string(val))
			octo[i] = append(octo[i], num)
		}
	}
	newOcto := padMap(octo)

	for turn := 0; turn < turns; turn++ {
		flashes += execute(newOcto)
	}

	for y := 0; y < len(newOcto); y++ {
		for x := 0; x < len(newOcto[0]); x++ {
			fmt.Printf("%x ", newOcto[y][x])
		}
		fmt.Printf("\n")
	}

	return flashes
}

func part2(values []string) int64 {
	return 0
}

func execute(values [][]int) int {
	var flashes int = 0

	// Increase by 1
	for y := 1; y < len(values)-2; y++ {
		for x := 1; x < len(values[0])-2; x++ {
			values[y][x] += 1
		}
	}

	// Flash any octopus greater than 9 and increase any adjacent by 1

	for y := 1; y < len(values)-2; y++ {
		for x := 1; x < len(values[0])-2; x++ {
			flashes += flash(values, x, y)
		}
	}

	// Set any octopus that flashed to 0
	for y := 1; y < len(values)-2; y++ {
		for x := 1; x < len(values[0])-2; x++ {
			if values[y][x] > 9 && values[y][x] < 15 {
				values[y][x] = 0
			}
		}
	}

	return flashes
}

func flash(v [][]int, x int, y int) int {
	var flashes int = 0
	value := v[y][x]

	for y := 1; y < len(v)-2; y++ {
		for x := 1; x < len(v[0])-2; x++ {
			if v[y][x] > 9 && v[y][x] < 15 {
				if value > 9 && value < 15 {
					if v[y][x+1] < 9 {
						v[y][x+1] += 1
					}
					if v[y][x-1] < 9 {
						v[y][x-1] += 1
					}
					if v[y+1][x] < 9 {
						v[y+1][x] += 1
					}
					if v[y-1][x] < 9 {
						v[y-1][x] += 1
					}
					if v[y+1][x+1] < 9 {
						v[y+1][x+1] += 1
					}
					if v[y+1][x-1] < 9 {
						v[y+1][x-1] += 1
					}
					if v[y-1][x+1] < 9 {
						v[y-1][x+1] += 1
					}
					if v[y-1][x-1] < 9 {
						v[y-1][x-1] += 1
					}
				}
				flashes++
			}
		}
	}
	return flashes
}

func padMap(values [][]int) [][]int {
	var padValue = 0xf
	var topAndBottom []int
	for i := 0; i < len(values[0])+2; i++ {
		topAndBottom = append(topAndBottom, padValue)
	}
	newCaves := make([][]int, 0)
	newCaves = append(newCaves, [][]int{topAndBottom}...)
	for _, line := range values {
		temp := make([]int, padValue)
		temp = arrayAppend(line, padValue)
		newCaves = append(newCaves, [][]int{temp}...)
	}
	newCaves = append(newCaves, [][]int{topAndBottom}...)
	return newCaves
}

func arrayAppend(arr []int, padValue int) []int {
	arr = append(arr, padValue)
	arr = append([]int{padValue}, arr...)
	return arr
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
