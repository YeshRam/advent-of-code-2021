package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func sortString(input string) string {
	s := strings.Split(input, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func contains(a string, b string) bool {
	if a == "" || b == "" {
		return false
	}

	for _, r := range b {
		if !strings.ContainsRune(a, r) {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open("day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	totalSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " | ")
		signal := strings.Split(split[0], " ")
		output := strings.Split(split[1], " ")

		decoder := map[string]int{}
		encoder := map[int]string{}

		for len(decoder) < 10 { // Keep looping through signals until we've found all encodings
			for _, val := range signal {
				sorted := sortString(val)
				length := len(sorted)

				if length == 2 {
					// Digit is 1
					decoder[sorted] = 1
					encoder[1] = sorted
				} else if length == 3 {
					// Digit is 7
					decoder[sorted] = 7
					encoder[7] = sorted
				} else if length == 4 {
					// Digit is 4
					decoder[sorted] = 4
					encoder[4] = sorted
				} else if length == 7 {
					// Digit is 8
					decoder[sorted] = 8
					encoder[8] = sorted
				} else if length == 5 {
					// Digit is 2, 3, 5

					_, found6 := encoder[6]
					_, found7 := encoder[7]

					if !found6 || !found7 {
						continue // We don't have enough info yet
					}

					if contains(sorted, encoder[7]) {
						// Digit is 3
						decoder[sorted] = 3
						encoder[3] = sorted
					} else if contains(encoder[6], sorted) {
						// Digit is 5
						decoder[sorted] = 5
						encoder[5] = sorted
					} else {
						// Digit is 2
						decoder[sorted] = 2
						encoder[2] = sorted
					}
				} else if length == 6 {
					// Digit is 0, 6, or 9

					_, found4 := encoder[4]
					_, found7 := encoder[7]

					if !found4 || !found7 {
						continue // We don't have enough info yet
					}

					if contains(sorted, encoder[4]) {
						// Digit is 9
						decoder[sorted] = 9
						encoder[9] = sorted
					} else if contains(sorted, encoder[7]) {
						// Digit is 0
						decoder[sorted] = 0
						encoder[0] = sorted
					} else {
						// Digit is 6
						decoder[sorted] = 6
						encoder[6] = sorted
					}
				}
			}
		}

		outputSum := 0
		for i, val := range output {
			sorted := sortString(val)
			outputSum += int(math.Pow10(len(output)-1-i)) * decoder[sorted]
			fmt.Printf("%d\t", decoder[sorted])
		}
		fmt.Printf("Output: %d\n", outputSum)
		totalSum += outputSum
	}

	fmt.Printf("\nTotal Sum: %d\n", totalSum)
}
