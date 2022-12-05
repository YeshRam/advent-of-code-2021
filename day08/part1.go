package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " | ")
		output := strings.Split(split[1], " ")

		for _, val := range output {
			length := len(val)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				count++
			}
		}
	}

	fmt.Println(count)
}
