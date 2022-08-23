package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	lines, err := readLines("Day 9\\input.txt")
	if err != nil {
		fmt.Println("Error parsing value")
	}
	fmt.Printf("Part1: %d\n", part1(lines))
	fmt.Printf("Part2: %d\n", part2(lines))
}

func part1(values []string) int {
	var caves [][]int = make([][]int, len(values))
	for i, line := range values {
		for _, val := range line {
			num, _ := strconv.Atoi(string(val))
			caves[i] = append(caves[i], num)
		}
	}
	newCaves := padCaves(caves)
	lows := findLows(newCaves)
	return calculateTotal(lows)
}

func part2(values []string) int {
	var caves [][]int = make([][]int, len(values))
	for i, line := range values {
		for _, val := range line {
			num, _ := strconv.Atoi(string(val))
			caves[i] = append(caves[i], num)
		}
	}
	newCaves := padCaves(caves)
	lowPoints := findLowCoords(newCaves)
	floodedAreas := flood(newCaves, lowPoints)
	sort.Ints(floodedAreas[:])
	end := len(floodedAreas) - 1
	return floodedAreas[end] * floodedAreas[end-1] * floodedAreas[end-2]
}

func flood(terrain [][]int, points []Point) []int {
	var counts []int
	var count int
	lastCount := 0
	for _, point := range points {
		terrain = fill(terrain, point)
		count = find10(terrain)
		counts = append(counts, count-lastCount)
		lastCount = count
	}
	return counts
}

func find10(terrain [][]int) int {
	var count int = 0
	for x := 1; x < len(terrain)-1; x++ {
		for y := 1; y < len(terrain[0])-1; y++ {
			if terrain[x][y] == 10 {
				count++
			}
		}
	}
	return count
}

func fill(terrain [][]int, point Point) [][]int {
	if point.x < 0 || point.y < 0 || point.x >= len(terrain[0]) || point.y >= len(terrain) {
		return terrain
	}
	if terrain[point.y][point.x] >= 9 {
		return terrain
	}
	terrain[point.y][point.x] = 10
	terrain = fill(terrain, Point{point.x - 1, point.y})
	terrain = fill(terrain, Point{point.x + 1, point.y})
	terrain = fill(terrain, Point{point.x, point.y - 1})
	terrain = fill(terrain, Point{point.x, point.y + 1})
	return terrain
}

func findLows(v [][]int) []int {
	var lowPoints []int
	for x := 1; x < len(v[0])-1; x++ {
		for y := 1; y < len(v)-1; y++ {
			h := v[y][x]
			if h < v[y+1][x] && h < v[y][x+1] && h < v[y-1][x] && h < v[y][x-1] {
				lowPoints = append(lowPoints, h)
			}
		}
	}
	return lowPoints
}

func findLowCoords(v [][]int) []Point {
	var lowPoints []Point
	for x := 1; x < len(v[0])-1; x++ {
		for y := 1; y < len(v)-1; y++ {
			h := v[y][x]
			if h < v[y+1][x] && h < v[y][x+1] && h < v[y-1][x] && h < v[y][x-1] {
				lowPoints = append(lowPoints, Point{x, y})
			}
		}
	}
	return lowPoints
}

func padCaves(values [][]int) [][]int {
	var topAndBottom []int
	for i := 0; i < len(values[0])+2; i++ {
		topAndBottom = append(topAndBottom, 9)
	}
	newCaves := make([][]int, 0)
	newCaves = append(newCaves, [][]int{topAndBottom}...)
	for _, line := range values {
		temp := make([]int, 9)
		temp = arrayAppend(line)
		newCaves = append(newCaves, [][]int{temp}...)
	}
	newCaves = append(newCaves, [][]int{topAndBottom}...)
	return newCaves
}

func arrayAppend(arr []int) []int {
	arr = append(arr, 9)
	arr = append([]int{9}, arr...)
	return arr
}

func calculateTotal(lows []int) int {
	total := 0
	for _, n := range lows {
		total += n + 1
	}
	return total
}
