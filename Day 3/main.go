package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	lines, err := readLines("Day 3\\input.txt")
	if err != nil {
		fmt.Println("Error parsing value")
	}

	fmt.Printf("Part1: %d\n", part1(lines))
	fmt.Printf("Part2: %d\n", part2(lines))
}

func part1(lines []string) int64 {
	var values = make([]int, len(lines[0]))
	for _, line := range lines {
		for i, elem := range line {
			if elem == '1' {
				values[i] += 1
			}
		}
	}
	var buff = make([]byte, len(lines[0]))
	for i, val := range values {
		if val > (len(lines) / 2) {
			buff[i] = '1'
		} else {
			buff[i] = '0'
		}
	}
	gamma, _ := strconv.ParseInt(string(buff[:]), 2, 64)
	epsilon, _ := strconv.ParseInt(invert(string(buff[:])), 2, 64)
	result := epsilon * gamma
	fmt.Printf("Gamma: %d Epsilon: %d \n", gamma, epsilon)
	return result
}

func invert(input string) string {
	var output = make([]byte, len(input))
	for i, ch := range input {
		if ch == '1' {
			output[i] = '0'
		} else {
			output[i] = '1'
		}
	}
	return string(output[:])
}

func part2(lines []string) int64 {
	var ox int64
	var co int64
	list1 := lines
	list2 := lines
	var out1 []string
	var out2 []string
	var out3 []string
	var out4 []string
	for i, _ := range lines[0] {
		if len(list1) > 1 {
			for _, x := range list1 {
				if x[i] == '0' {
					out1 = append(out1, x)
				}
			}
			for _, x := range list1 {
				if x[i] == '1' {
					out2 = append(out2, x)
				}
			}
		}
		if len(list2) > 1 {
			for _, y := range list2 {
				if y[i] == '1' {
					out3 = append(out3, y)
				}
			}
			for _, y := range list2 {
				if y[i] == '0' {
					out4 = append(out4, y)
				}
			}
		}
		if len(out1) > len(out2) {
			list1 = out1
		} else {
			list1 = out2
		}
		if len(out3) < len(out4) {
			list2 = out3
		} else {
			list2 = out4
		}

		if len(list1) == 1 {
			ox, _ = strconv.ParseInt(list1[0], 2, 64)
		}
		if len(list2) == 1 {
			co, _ = strconv.ParseInt(list2[0], 2, 64)
		}

		out1 = []string{}
		out2 = []string{}
		out3 = []string{}
		out4 = []string{}
	}
	return ox * co
}
