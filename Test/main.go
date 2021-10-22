// Code for Advent of Code 2018 Day 1

package main

import (
	"bufio"
	"fmt"
	"log"
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
	var position int64 = 0
	var duplicate bool = false
	var firstrun bool = true
	var freqs []int64
	var freq int64 = 0
	var iterations int64 = 0
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for {
		for _, line := range lines {
			var operator = line[0]
			var value = line[1:]
			val, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				fmt.Println("Error parsing value")
			}
			if operator == '+' {
				position = position + val
			}
			if operator == '-' {
				position = position - val
			}
			if !duplicate && contains(freqs, position) {
				freq = position
				duplicate = true
			}
			freqs = append(freqs, position)
		}
		if duplicate {
			break
		}
		if firstrun {
			fmt.Printf("\nPart 1 Result: %d", position)
			firstrun = false
		}
		iterations++
	}
	fmt.Printf("\nPart 2 Result: %d after %d iterations\n", freq, iterations)
}

func contains(arr []int64, val int64) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}
