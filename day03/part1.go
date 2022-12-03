package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const BinaryLength = 12

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	oneCounts := [BinaryLength]int{}
	numLines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i, bit := range line {
			if bit == '1' {
				oneCounts[i]++
			}
		}
		numLines++
	}

	gammaRate := [BinaryLength]int{}
	epsilonRate := [BinaryLength]int{}
	for i := 0; i < BinaryLength; i++ {
		oneCount := oneCounts[i]
		zeroCount := numLines - oneCount
		if oneCount > zeroCount {
			gammaRate[i] = 1
			epsilonRate[i] = 0
		} else {
			gammaRate[i] = 0
			epsilonRate[i] = 1
		}
	}

	powerConsumption := binaryToDecimal(gammaRate) * binaryToDecimal(epsilonRate)
	fmt.Println(powerConsumption)
}

func binaryToDecimal(input [BinaryLength]int) int {
	result := 0
	power := 0
	for i := len(input) - 1; i >= 0; i-- {
		result += int(math.Pow(2, float64(power))) * input[i]
		power++
	}
	return result
}
