package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var count = 0
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	prev, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if depth > prev {
			count++
		}

		prev = depth
	}

	fmt.Println(count)
}
