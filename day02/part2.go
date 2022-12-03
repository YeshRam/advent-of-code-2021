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
	aim := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		direction := split[0]
		units, _ := strconv.Atoi(split[1])

		if direction == "up" {
			aim -= units
		} else if direction == "down" {
			aim += units
		} else if direction == "forward" {
			x += units
			y += units * aim
			if y < 0 {
				y = 0
			}
		}
	}

	fmt.Println(x * y)
}
