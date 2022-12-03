package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	x := 0
	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		direction := split[0]
		distance, _ := strconv.Atoi(split[1])

		if direction == "forward" {
			x += distance
		} else if direction == "up" {
			y -= distance
			if y < 0 {
				y = 0
			}
		} else if direction == "down" {
			y += distance
		}
	}

	fmt.Println(x * y)

}
