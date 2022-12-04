package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	positions := []int{}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	split := strings.Split(line, ",")
	for _, val := range split {
		val, _ := strconv.Atoi(val)
		positions = append(positions, val)
	}

	median := -1
	sort.Ints(positions) // Sorting here but a more efficient way to find a median is with two heaps
	if len(positions)%2 == 0 {
		median = (positions[len(positions)/2] + positions[len(positions)/2-1]) / 2
	} else {
		median = positions[len(positions)/2]
	}

	fuel := 0
	for _, val := range positions {
		fuel += int(math.Abs(float64(val - median)))
	}

	fmt.Printf("Position: %d\n", median)
	fmt.Printf("Fuel needed: %d\n", fuel)
}
