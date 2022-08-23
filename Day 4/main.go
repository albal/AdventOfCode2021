package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, [][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var lines []string

	var numbers []string
	var boards [][]string
	scanner := bufio.NewScanner(file)
	scanner.Split(crunchSplitFunc)
	space := regexp.MustCompile("\\s+")
	for scanner.Scan() {
		lines = append(lines, space.ReplaceAllString(scanner.Text(), " "))
	}

	numbers = strings.Split(lines[0], ",")

	for _, board := range lines[1:] {
		boards = append(boards, strings.Split(board[1:], " "))
	}

	return numbers, boards, scanner.Err()
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
	numbers, boards, err := readLines("Day 4\\input.txt")
	if err != nil {
		fmt.Println("Error parsing value")
	}

	fmt.Printf("Part1: %d\n", part1(numbers, boards))
	fmt.Printf("Part2: %d\n", part2(numbers, boards))

	numbers = []string{"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1"}
	boards = [][]string{[]string{"22", "13", "17", "11", "0", "8", "2", "23", "4", "24", "21", "9", "14", "16", "7", "6", "10", "3", "18", "5", "1", "12", "20", "15", "19"}, []string{"3", "15", "0", "2", "22", "9", "18", "13", "17", "5", "19", "8", "7", "25", "23", "20", "11", "10","24", "4", "14", "21", "16", "12", "6"}, []string{"14", "21", "17", "24", "4", "10", "16", "15", "9", "19", "18", "8", "23", "26", "20", "22", "11", "13", "6", "5", "2", "0", "12", "3", "7"}})


	fmt.Printf("PartX %d\n", part1(numbers, boards))
}

func part1(numbers []string, boards [][]string) int {
	var result []string
	var boardNumber int
	var lastNumber string
	for _, n := range numbers {
		for i, board := range boards {
			for j, number := range board {
				if n == number {
					boards[i][j] = "M" + n
				}
			}
		}
		result = checkBoards(boards)
		if len(result) != 0 {
			fmt.Printf("We found a winner and it is board %s with number %s\n", result, n)
			boardNumber, _ = strconv.Atoi(result[0])
			lastNumber = n
			fmt.Printf("%s\n", boards[boardNumber])
			break
		}
	}
	// Calculate winning board
	return calculateAnswer(boards[boardNumber], lastNumber)
}

func part2(numbers []string, boards [][]string) int {
	var result []string
	var lastBoard []string
	var res string
	for _, n := range numbers {
		for i, board := range boards {
			for j, number := range board {
				if n == number {
					boards[i][j] = "M" + n
				}
			}
		}
		result = checkBoards(boards)
		if len(result) != 0 {
			for i := range result {
				br, _ := strconv.Atoi(result[len(result) - 1 -i])
				boards = removeBoard(boards, br)
			}
		}
		if len(boards) == 1 && lastBoard == nil {
			lastBoard = boards[0]
			fmt.Printf("Board: %s\n", boards[0])
		}
		if len(boards) == 0 {
			fmt.Printf("We have the last board with number %s\n", n)
			res = n
			break
		}
	}
	// Calculate winning board
	return calculateAnswer(lastBoard, res)
}

func calculateAnswer(board []string, number string) int {
	var total int = 0
	num, _ := strconv.Atoi(number)
	for _, val := range board {
		if !strings.Contains(val, "M") {
			if val != number {
				value, _ := strconv.Atoi(val)
				total = total + value
			}
		}
	}
	return total * num
}

func makeRange (num int) []string {
	var intRange []string
	for i := 0; i < num; i++ {
		val := strconv.Itoa(i + 1)
		intRange = append(intRange, val)
	}
	return intRange
}

func removeBoard(s [][]string, r int) [][]string {
	firstBoards := s[:r]
	lastBoards := s[r+1:]
	return append(firstBoards, lastBoards...)
}

func checkBoards(boards [][]string) []string {
	var boardsWon []string
	for n, board := range boards {
		if rowsMatch(board) || colsMatch(board) {
			boardsWon = append(boardsWon, strconv.Itoa(n))
		}
	}
	return boardsWon
}

func rowsMatch(b []string) bool {

	if m(b[0]) && m(b[1]) && m(b[2]) && m(b[3]) && m(b[4]) {
		return true
	}
	if m(b[5]) && m(b[6]) && m(b[7]) && m(b[8]) && m(b[9]) {
		return true
	}
	if m(b[10]) && m(b[11]) && m(b[12]) && m(b[13]) && m(b[14]) {
		return true
	}
	if m(b[15]) && m(b[16]) && m(b[17]) && m(b[18]) && m(b[19]) {
		return true
	}
	if m(b[20]) && m(b[21]) && m(b[22]) && m(b[23]) && m(b[24]) {
		return true
	}
	return false
}

func colsMatch(b []string) bool {
	if m(b[0]) && m(b[5]) && m(b[10]) && m(b[15]) && m(b[20]) {
		return true
	}
	if m(b[1]) && m(b[6]) && m(b[11]) && m(b[16]) && m(b[21]) {
		return true
	}
	if m(b[2]) && m(b[7]) && m(b[12]) && m(b[17]) && m(b[22]) {
		return true
	}
	if m(b[3]) && m(b[8]) && m(b[13]) && m(b[18]) && m(b[23]) {
		return true
	}
	if m(b[4]) && m(b[9]) && m(b[14]) && m(b[19]) && m(b[24]) {
		return true
	}
	return false
}

func m(item string) bool {
	if strings.Contains(item, "M") {
		return true
	}
	return false
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
