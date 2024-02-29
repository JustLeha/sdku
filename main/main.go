package main

import (
	"fmt"
	"os"
)

func printStr(s string) {
	fmt.Print(s)
}

func validateInput(args []string) ([9][9]int, bool) {
	var board [9][9]int
	if len(args) != 9 {
		return board, false
	}
	for i, row := range args {
		if len(row) != 9 {
			return board, false
		}
		for j, ch := range row {
			if ch == '.' {
				board[i][j] = 0
			} else if ch >= '1' && ch <= '9' {
				board[i][j] = int(ch - '0')
			} else {
				return board, false
			}
		}
	}
	return board, true
}

func printBoard(board [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(board[i][j])
			if j != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func isSafe(board [9][9]int, row, col, num int) bool {
	for x := 0; x < 9; x++ {
		if board[row][x] == num || board[x][col] == num {
			return false
		}
	}
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

func solveSudoku(board [9][9]int) ([9][9]int, bool) {
	row, col, isEmpty := -1, -1, true
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				row, col, isEmpty = i, j, false
				break
			}
		}
		if !isEmpty {
			break
		}
	}
	if isEmpty {
		return board, true
	}
	for num := 1; num <= 9; num++ {
		if isSafe(board, row, col, num) {
			board[row][col] = num
			if solvedBoard, ok := solveSudoku(board); ok {
				return solvedBoard, true
			}
			board[row][col] = 0
		}
	}
	return board, false
}

func main() {
	args := os.Args[1:]
	board, isValid := validateInput(args)
	if !isValid {
		printStr("Error")
		fmt.Println()
		return
	}
	if solvedBoard, ok := solveSudoku(board); ok {
		printBoard(solvedBoard)
	} else {
		printStr("Error")
	}
}
