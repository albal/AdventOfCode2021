package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	lines, err := readLines("Day 8\\input.txt")
	if err != nil {
		fmt.Println("Error parsing value")
	}
	fmt.Printf("Part1: %d\n", part1(lines))
	fmt.Printf("Part2: %d\n", part2(lines))
}

func part1(values []string) int {
	total := 0
	for _, line := range values {
		toCheck := strings.Split(line, "|")[1]
		segments := strings.Split(toCheck, " ")
		for _, segment := range segments {
			switch segLen := len(segment); segLen {
			case 2:
				total += 1
			case 3:
				total += 1
			case 4:
				total += 1
			case 7:
				total += 1
			default:
			}
		}

	}
	return total
}

func part2(values []string) int64 {
	var total int64 = 0
	for _, line := range values {
		toCheck := strings.Split(line, "|")[0]
		segmentsUnordered := strings.Split(toCheck, " ")
		segments := sortValues(segmentsUnordered[:10])
		digits := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		for _, segment := range segments {
			segLen := len(segment)
			if segLen == 2 {
				digits[1] = segment
			} else if segLen == 3 {
				digits[7] = segment
			} else if segLen == 4 {
				digits[4] = segment
			} else if segLen == 5 {
				if isSegment(segment, digits[1]) { // 3 has 1 in it
					digits[3] = segment
				} else {
					var count = 0
					for _, c := range digits[4] {
						if strings.Contains(segment, string(c)) {
							count++
						}
					}
					if count == 3 {
						digits[5] = segment
					} else {
						digits[2] = segment
					}
				}
			} else if segLen == 6 {
				if isSegment(segment, digits[4]) {
					digits[9] = segment
				} else if isSegment(segment, digits[7]) {
					digits[0] = segment
				} else {
					digits[6] = segment
				}
			} else {
				digits[8] = segment
			}

		}
		total += calcDisplay(digits[:], strings.Trim(strings.Split(line, "|")[1], " "))
	}
	return total
}

func calcDisplay(digits []string, check string) int64 {
	output := ""
	vals := strings.Split(check, " ")
	for _, d := range vals {
		val := getDigit(digits, d)
		valStr := strconv.Itoa(int(val))
		output += valStr
	}
	number, _ := strconv.Atoi(output)
	return int64(number)
}

type byLength []string

func sortValues(list []string) []string {
	sort.Sort(byLength(list))
	return list
}

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func isSegment(segment string, value string) bool {
	var isPresent int = 0
	for _, c := range value {
		if strings.Contains(segment, string(c)) {
			isPresent++
		}
	}
	if isPresent == len(value) {
		return true
	}
	return false
}

func getDigit(digits []string, digit string) int {
	for n, check := range digits {
		if SortString(check) == SortString(digit) {
			return n
		}
	}
	return -1
}

func SortString(s string) string {
	r := strings.Split(s, "")
	sort.Strings(r)
	return strings.Join(r, "")
}
