package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const GridSize = 1000

type grid struct {
	grid [GridSize][GridSize]int
}

func newGrid(array [GridSize][GridSize]int) *grid {
	g := grid{grid: array}
	return &g
}

func printGrid(g *grid) {
	for i := 0; i < len(g.grid); i++ {
		for j := 0; j < len(g.grid[0]); j++ {
			fmt.Printf("%d\t", g.grid[j][i])
		}
		fmt.Println()
	}
	fmt.Println()
}

func plotLine(g *grid, x1 int, y1 int, x2 int, y2 int) {
	// Only considering horizontal lines
	if x1 == x2 {
		if y2 > y1 {
			for i := y1; i <= y2; i++ {
				g.grid[x1][i]++
			}
		} else {
			for i := y2; i <= y1; i++ {
				g.grid[x1][i]++
			}
		}
	} else if y1 == y2 {
		if x2 > x1 {
			for i := x1; i <= x2; i++ {
				g.grid[i][y1]++
			}
		} else {
			for i := x2; i <= x1; i++ {
				g.grid[i][y1]++
			}
		}
	}
}

func countOverlaps(g *grid) int {
	count := 0
	for i := 0; i < len(g.grid); i++ {
		for j := 0; j < len(g.grid[0]); j++ {
			if g.grid[i][j] >= 2 {
				count++
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	array := [GridSize][GridSize]int{}
	grid := newGrid(array)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " -> ")

		x1, _ := strconv.Atoi(strings.Split(split[0], ",")[0])
		y1, _ := strconv.Atoi(strings.Split(split[0], ",")[1])
		x2, _ := strconv.Atoi(strings.Split(split[1], ",")[0])
		y2, _ := strconv.Atoi(strings.Split(split[1], ",")[1])

		plotLine(grid, x1, y1, x2, y2)
	}

	//printGrid(grid)

	fmt.Println(countOverlaps(grid))
}
