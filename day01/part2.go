package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var scanner = bufio.NewScanner(file)
	var window [3]int
	var count = 0

	// Handle first three iterations to fill sliding window
	for i := 0; i < 3; i++ {
		scanner.Scan()
		window[i], _ = strconv.Atoi(scanner.Text())
	}

	for scanner.Scan() {
		prevSum := window[0] + window[1] + window[2]

		depth, _ := strconv.Atoi(scanner.Text())
		currSum := window[1] + window[2] + depth

		if currSum > prevSum {
			count++
		}

		window[0] = window[1]
		window[1] = window[2]
		window[2] = depth
	}

	fmt.Println(count)
}
