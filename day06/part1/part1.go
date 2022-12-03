package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const NumDays = 18

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
	//fmt.Printf("%d fish during initial state: %v\n", len(lanternfish), lanternfish)

	// Simulate
	timesSpawned := 0
	for day := 1; day <= NumDays; day++ {
		spawn := []int{}
		for i := range lanternfish {
			if lanternfish[i] == 0 {
				lanternfish[i] = 6
				spawn = append(spawn, 8)
				timesSpawned++
				continue
			}

			lanternfish[i]--
		}
		lanternfish = append(lanternfish, spawn...)
		//fmt.Printf("%d fish after %d days: %v\n", len(lanternfish), day, lanternfish)
	}

	fmt.Printf("%d fish after %d days\n", len(lanternfish), NumDays)
}
