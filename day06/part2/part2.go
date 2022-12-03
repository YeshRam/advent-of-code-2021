package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const NumDays = 256

func main() {
	file, err := os.Open("day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lanternfish := []int{}

	// Read input
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	split := strings.Split(line, ",")
	for _, val := range split {
		val, _ := strconv.Atoi(val)
		lanternfish = append(lanternfish, val)
	}

	// Simulate
	// Group lanternfish by time to spawn
	// Index: Days till spawn, Value: number of fish
	groups := [9]int{}

	for _, val := range lanternfish {
		groups[val]++
	}

	for day := 1; day <= NumDays; day++ {
		nextDay := [9]int{}
		nextDay[6] += groups[0]
		nextDay[8] += groups[0]
		for i := 1; i < 9; i++ {
			nextDay[i-1] += groups[i]
		}
		groups = nextDay
	}

	total := 0
	for _, val := range groups {
		total += val
	}
	fmt.Printf("%d fish after %d days\n", total, NumDays)
}
