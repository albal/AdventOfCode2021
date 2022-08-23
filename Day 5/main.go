package main

import (
	"bufio"
	"fmt"
	"os"
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
	lines, err := readLines("Day 5\\input.txt")
	if err != nil {
		fmt.Println("Error parsing value")
	}

	fmt.Printf("Part1: %d\n", part1(lines))
	fmt.Printf("Part2: %d\n", part2(lines))
}

func part1(lines []string) int {
	var board [1000][1000]int
	for n := 0; n < len(board); n++ {
		board[n] = [len(board)]int{}
	}
	for _, line := range lines {
		var (
			x1 int
			y1 int
			x2 int
			y2 int
		)
		r := strings.NewReader(line)
		_, _ = fmt.Fscanf(r, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		if x1 == x2 {
			len := max(y1, y2) - min(y1, y2) + 1
			offset := min(y1, y2)
			for i := 0; i < len; i++ {
				board[x1][i+offset] = board[x1][i+offset] + 1
			}
		}
		if y1 == y2 {
			len := max(x1, x2) - min(x1, x2) + 1
			offset := min(x1, x2)
			for i := 0; i < len; i++ {
				board[i+offset][y1] = board[i+offset][y1] + 1
			}
		}
	}

	maxVal := 2 // Find where lines cross more than once

	var count int = 0
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if board[x][y] >= maxVal {
				count++
			}
		}
	}
	return count
}

func part2(lines []string) int {
	var board [1000][1000]int
	for n := 0; n < len(board); n++ {
		board[n] = [len(board)]int{}
	}
	for _, line := range lines {
		var (
			x1 int
			y1 int
			x2 int
			y2 int
		)
		r := strings.NewReader(line)
		_, _ = fmt.Fscanf(r, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		if x1 == x2 {
			length := max(y1, y2) - min(y1, y2) + 1
			offset := min(y1, y2)
			for i := 0; i < length; i++ {
				board[x1][i+offset] += 1
			}
		}

		if y1 == y2 {
			length := max(x1, x2) - min(x1, x2) + 1
			offset := min(x1, x2)
			for i := 0; i < length; i++ {
				board[i+offset][y1] += 1
			}
		}

		if x1 != x2 && y1 != y2 {
			length := max(x1, x2) - min(x1, x2) + 1
			if x1 < x2 && y1 < y2 {
				for j := 0; j < length; j++ {
					board[x1+j][y1+j] += 1
				}
			} else if x1 < x2 && y1 > y2 {
				for j := 0; j < length; j++ {
					board[x1+j][y1-j] += 1
				}
			} else if x1 > x2 && y1 > y2 {
				for j := 0; j < length; j++ {
					board[x1-j][y1-j] += 1
				}
			} else if x1 > x2 && y1 < y2 {
				for j := 0; j < length; j++ {
					board[x1-j][y1+j] += 1
				}
			}
		}
	}

	maxVal := 2 // Find where lines cross more than once

	var count int = 0
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if board[x][y] >= maxVal {
				count++
			}
		}
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
