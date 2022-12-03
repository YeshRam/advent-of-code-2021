package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const NumLines = 1000
const BinaryLength = 12

func main() {
	file, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]string, NumLines)
	o2 := map[string]bool{}
	co2 := map[string]bool{}
	oneCounts := [BinaryLength]int{}

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, bit := range line {
			if bit == '1' {
				oneCounts[i]++
			}
		}
		o2[line] = true
		co2[line] = true
		data[row] = line
		row++
	}

	for col := 0; col < len(data[0]); col++ {
		oneCount := 0
		zeroCount := 0
		for row := range o2 {
			if row[col] == '1' {
				oneCount++
			} else {
				zeroCount++
			}
		}

		for row := 0; row < len(data); row++ {
			if oneCount >= zeroCount {
				if data[row][col] == '0' {
					delete(o2, data[row])
				}
			} else {
				if data[row][col] == '1' {
					delete(o2, data[row])
				}
			}
		}

		if len(o2) == 1 {
			break
		}
	}

	generatorRating := 0
	for key := range o2 {
		generatorRating = binaryToDecimal(key)
	}

	for col := 0; col < len(data[0]); col++ {
		oneCount := 0
		zeroCount := 0
		for row := range co2 {
			if row[col] == '1' {
				oneCount++
			} else {
				zeroCount++
			}
		}

		for row := 0; row < len(data); row++ {
			if oneCount >= zeroCount {
				if data[row][col] == '1' {
					delete(co2, data[row])
				}
			} else {
				if data[row][col] == '0' {
					delete(co2, data[row])
				}
			}
		}

		if len(co2) == 1 {
			break
		}
	}

	scrubberRating := 0
	for key := range co2 {
		scrubberRating = binaryToDecimal(key)
	}

	lifeSupportRating := generatorRating * scrubberRating
	fmt.Println(lifeSupportRating)
}

func binaryToDecimal(input string) int {
	runes := []rune(input)

	result := 0
	power := 0
	for i := len(runes) - 1; i >= 0; i-- {
		result += int(math.Pow(2, float64(power))) * int(runes[i]-'0')
		power++
	}
	return result
}
