package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

// clear the console on every render cycle to create a frame changing illusion
func Clear_Console() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Random_State(width, height int) [][]int {
	result := [][]int{}
	random_threshold := 0.3
	for i := 0; i < height; i++ {
		temp := make([]int, width)
		for idx := range temp {
			if rand.Float64() >= random_threshold {
				temp[idx] = 1
			} else {
				temp[idx] = 0
			}
		}
		result = append(result, temp)
	}
	return result
}

func Render(random_state [][]int) {
	result := [][]string{}
	for _, elem := range random_state {
		row := []string{}
		for i := 0; i < len(elem); i++ {
			// if it's dead literally do nothing
			if elem[i] == 0 {
				row = append(row, " ")
			} else if elem[i] == 1 {
				row = append(row, "*")
			}
		}
		result = append(result, row)
	}

	Clear_Console()
	for _, elem := range result {
		fmt.Println(strings.Join(elem, " "))
	}
}

/*
assumptions:
 1. All matrices will be square matrices
*/
func Next_Board_State(initial_board [][]int) [][]int {
	rows := len(initial_board)
	cols := len(initial_board[0])
	result := make([][]int, rows)
	copy(result, initial_board)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// create temporary array to store values for cell state computation
			temp_arr := []int{}
			var a int
			var b int
			var c int
			var d int
			var e int
			var f int
			var g int
			var h int

			// Regular top edge with 5 neighbors
			if i == 0 && (j != 0 && j != rows-1) {
				// top edge with 5 neighbors
				a, b, c, d, e = initial_board[i][j-1], initial_board[i+1][j-1], initial_board[i+1][j], initial_board[i+1][j+1], initial_board[i][j+1]
				// Regular bottom edge with 5 neighbors
			} else if i == rows-1 && (j != 0 && j != rows-1) {
				a, b, c, d, e = initial_board[i][j-1], initial_board[i-1][j-1], initial_board[i-1][j], initial_board[i-1][j+1], initial_board[i][j+1]
				// Regular right edge with 5 neighbors
			} else if j == rows-1 && (i != 0 && i != rows-1) {
				a, b, c, d, e = initial_board[i-1][j], initial_board[i-1][j-1], initial_board[i][j-1], initial_board[i+1][j-1], initial_board[i+1][j]
				// Regular left edge with 5 neighbors
			} else if j == 0 && (i != 0 && i != rows-1) {
				a, b, c, d, e = initial_board[i-1][j], initial_board[i-1][j+1], initial_board[i][j+1], initial_board[i+1][j+1], initial_board[i+1][j]
				// Top Left Corner
			} else if i == 0 && j == 0 {
				a, b, c = initial_board[i][j+1], initial_board[i+1][j+1], initial_board[i+1][j]
				// Top Right Corner
			} else if i == 0 && j == rows-1 {
				a, b, c = initial_board[i][j-1], initial_board[i+1][j-1], initial_board[i+1][j]
				// Bottom Left Corner
			} else if i == rows-1 && j == 0 {
				a, b, c = initial_board[i-1][j], initial_board[i-1][j+1], initial_board[i][j+1]
				// Bottom Right Corner
			} else if i == rows-1 && j == rows-1 {
				a, b, c = initial_board[i-1][j], initial_board[i-1][j-1], initial_board[i][j-1]
				// Trivial case where the cell lies in the middle with 8 neighbors
			} else {
				a, b, c, d, e, f, g, h = initial_board[i-1][j-1], initial_board[i-1][j], initial_board[i-1][j+1], initial_board[i][j+1], initial_board[i+1][j+1], initial_board[i+1][j], initial_board[i+1][j-1], initial_board[i][j-1]
			}
			temp_arr = append(temp_arr, a, b, c, d, e, f, g, h)

			// calculate the rules
			count := 0
			for _, elem := range temp_arr {
				count += elem
			}
			// fmt.Printf("The number of live cells is %d\n", count)

			// implement the rules
			if initial_board[i][j] == 1 && (count == 0 || count == 1) {
				result[i][j] = 0
			} else if initial_board[i][j] == 1 && count == 2 || count == 3 {
				result[i][j] = 1
			} else if initial_board[i][j] == 1 && count > 3 {
				result[i][j] = 0
			} else if (initial_board[i][j] == 0) && (count == 3) {
				result[i][j] = 1
			}
		}
	}
	return result
}

func main() {

	starting_state := Random_State(20, 20)
	next_state := Next_Board_State(starting_state)

	for {
		Render(next_state)
		time.Sleep(600 * time.Millisecond)
		next_state = Next_Board_State(next_state)
	}
}
