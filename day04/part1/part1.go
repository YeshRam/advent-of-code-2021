package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bingo struct {
	board  [5][5]int
	marked [5][5]bool
}

func newBoard(array [5][5]int) *bingo {
	b := bingo{board: array, marked: [5][5]bool{}}
	return &b
}

func printBoard(b *bingo) {
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[0]); j++ {
			fmt.Printf("%d\t", b.board[i][j])
		}
		fmt.Print("\t|\t")
		for j := 0; j < len(b.board[0]); j++ {
			fmt.Printf("%t\t", b.marked[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func markBoard(b *bingo, called int) {
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[0]); j++ {
			if b.board[i][j] == called {
				b.marked[i][j] = true
			}
		}
	}
}

func checkWin(b *bingo) bool {
	// Check rows for win
	for i := 0; i < len(b.board); i++ {
		count := 0
		for j := 0; j < len(b.board[0]); j++ {
			if b.marked[i][j] == true {
				count++
			}

			if count == 5 {
				return true
			}
		}
	}

	// Check cols for win
	for j := 0; j < len(b.board[0]); j++ {
		count := 0
		for i := 0; i < len(b.board); i++ {
			if b.marked[i][j] == true {
				count++
			}

			if count == 5 {
				return true
			}
		}
	}

	return false
}

func unmarkedSum(b *bingo) int {
	sum := 0
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[0]); j++ {
			if b.marked[i][j] == false {
				sum += b.board[i][j]
			}
		}
	}
	return sum
}

func main() {
	file, err := os.Open("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	split := strings.Split(scanner.Text(), ",")
	numbers := []int{}
	for _, v := range split {
		number, _ := strconv.Atoi(v)
		numbers = append(numbers, number)
	}

	scanner.Scan() // Skip the next empty line

	var boards []*bingo
	var x = 0
	var arr [5][5]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			boards = append(boards, newBoard(arr))
			x = 0
		} else {
			row := strings.Fields(line)
			for y, val := range row {
				val, _ := strconv.Atoi(val)
				arr[x][y] = val
			}
			x++
		}
	}
	boards = append(boards, newBoard(arr)) // scanner.Scan() doesn't process a newline at the end of the file

	result := -1
out:
	for i := 0; i < len(numbers); i++ {
		for _, b := range boards {
			markBoard(b, numbers[i])

			if checkWin(b) {
				result = unmarkedSum(b) * numbers[i]
				printBoard(b)
				break out
			}
		}
	}

	fmt.Println(result)
}
