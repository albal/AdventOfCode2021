package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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

func crunchSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := strings.Index(string(data), "\r\n\r\n"); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return
}

func main() {
	lines, err := readLines("Day 7\\input.txt")
	if err != nil {
		fmt.Println("Error parsing value")
	}
	strValues := strings.Split(lines[0], ",")
	values := make([]int64, 0, len(strValues))
	for _, value := range strValues {
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			log.Print(err)
			continue
		}
		values = append(values, v)
	}
	fmt.Printf("Part1: %d\n", part1(values))

	values = make([]int64, 0, len(strValues))
	for _, value := range strValues {
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			log.Print(err)
			continue
		}
		values = append(values, v)
	}

	fmt.Printf("Part2: %d\n", part2(values))
}

func part1(values []int64) int64 {
	var fuel int64 = 0
	var fuels []int64
	for _, n := range values {
		fuel = 0
		for _, p := range values {
			fuel += abs(n - p)
		}
		fuels = append(fuels, fuel)
	}
	return minFuel(fuels)
}

func part2(values []int64) int64 {
	//values = []int64{16,1,2,0,4,2,7,1,2,14}
	var fuels []int64
	for _, n := range values {
		var fuel int64 = 0
		for _, p := range values {
			var fuelToAdd = move(n, p)
			fuel += fuelToAdd
		}
		fuels = append(fuels, fuel)
	}
	return minFuel(fuels)
}

func move(n int64, p int64) int64 {
	// Triangle numbers
	pos1 := min(n, p)
	pos2 := max(n, p)
	len := pos2 - pos1
	return len * (len + 1) / 2
}

func minFuel(values []int64) int64 {
	var min int64 = math.MaxInt64
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func abs(value int64) int64 {
	if value < 0 {
		return value * -1
	}
	return value
}
