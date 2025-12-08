package day4

import (
	"fmt"
	"strings"
)

func Part1(input []byte) (string, error) {
	grid := make([][]rune, 0)
	for l := range strings.SplitSeq(string(input), "\n") {
		grid = append(grid, []rune(l))
	}
	adjacent := adjacencyMatrix(grid)
	var sln int
	for row := range adjacent {
		for col := range adjacent[row] {
			if adjacent[row][col] < 4 &&
				grid[row][col] == '@' {
				sln++
			}
		}
	}
	return fmt.Sprint(sln), nil
}

func Part2(input []byte) (string, error) {
	grid := make([][]rune, 0)
	for l := range strings.SplitSeq(string(input), "\n") {
		grid = append(grid, []rune(l))
	}
	var sln int
	for {
		adjacent := adjacencyMatrix(grid)
		var reachable int
		for row := range adjacent {
			for col := range adjacent[row] {
				if adjacent[row][col] < 4 &&
					grid[row][col] == '@' {
					reachable++
					grid[row][col] = '.'
				}
			}
		}
		if reachable == 0 {
			break
		}
		sln += reachable
	}
	return fmt.Sprint(sln), nil
}

func adjacencyMatrix(grid [][]rune) [][]int {
	adjacent := make([][]int, len(grid))
	for i := range len(adjacent) {
		adjacent[i] = make([]int, len(grid[0]))
	}
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '.' {
				continue
			}
			// up
			if row > 0 &&
				grid[row-1][col] == '@' {
				adjacent[row][col]++
			}
			// down
			if row < len(grid)-1 &&
				grid[row+1][col] == '@' {
				adjacent[row][col]++
			}
			// left
			if col > 0 &&
				grid[row][col-1] == '@' {
				adjacent[row][col]++
			}
			// right
			if col < len(grid[row])-1 &&
				grid[row][col+1] == '@' {
				adjacent[row][col]++
			}
			// up/left
			if row > 0 &&
				col > 0 &&
				grid[row-1][col-1] == '@' {
				adjacent[row][col]++
			}
			// up/right
			if row > 0 &&
				col < len(grid[row])-1 &&
				grid[row-1][col+1] == '@' {
				adjacent[row][col]++
			}
			// down/left
			if row < len(grid)-1 &&
				col > 0 &&
				grid[row+1][col-1] == '@' {
				adjacent[row][col]++
			}
			// down/right
			if row < len(grid)-1 &&
				col < len(grid[row])-1 &&
				grid[row+1][col+1] == '@' {
				adjacent[row][col]++
			}
		}
	}
	return adjacent
}
