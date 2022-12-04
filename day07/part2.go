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

func main() {
	file, err := os.Open("day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	positions := []int{}

	// Read input
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	split := strings.Split(line, ",")
	for _, val := range split {
		val, _ := strconv.Atoi(val)
		positions = append(positions, val)
	}

	// Precalculate fuel it takes based on number of steps
	// Index: number of steps, Value: fuel needed
	fuel := [2000]int{0, 1}
	for i := 2; i < 2000; i++ {
		fuel[i] = fuel[i-1] + i
	}

	// Find min and max positions. Our ideal position must be somewhere in between.
	minPos := math.MaxInt
	maxPos := math.MinInt
	for _, val := range positions {
		if val > maxPos {
			maxPos = val
		} else if val < minPos {
			minPos = val
		}
	}

	// Try every possible position and determine the position which requires the least total fuel.
	minFuel := math.MaxInt
	idealPos := 0
	for i := minPos; i <= maxPos; i++ {
		totalFuel := 0
		for _, val := range positions {
			totalFuel += fuel[int(math.Abs(float64(i-val)))]
		}

		if totalFuel < minFuel {
			minFuel = totalFuel
			idealPos = i
		}
	}

	fmt.Printf("We should align at position %d using %d fuel.\n", idealPos, minFuel)
}
