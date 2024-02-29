package golang

// updatePossibilities creates and returns a 3D slice containing possible values for each cell

func UpdatePossibilities(board [9][9]int) [9][9][]int {
	possibilities := [9][9][]int{}

	// Initialize all possibilities with numbers 1-9
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				possibilities[i][j] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			}
		}
	}

	// Update possibilities based on existing numbers on the board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != 0 {
				num := board[i][j]
				// Remove 'num' from possibilities in the same row, column, and box
				for k := 0; k < 9; k++ {
					// Row
					possibilities[i][k] = remove(possibilities[i][k], num)
					// Column
					possibilities[k][j] = remove(possibilities[k][j], num)
					// 3x3 box
					boxRow, boxCol := i/3*3, j/3*3
					for br := 0; br < 3; br++ {
						for bc := 0; bc < 3; bc++ {
							possibilities[boxRow+br][boxCol+bc] = remove(possibilities[boxRow+br][boxCol+bc], num)
						}
					}
				}
			}
		}
	}

	return possibilities
}

// remove is a helper function to remove a number from a slice
func remove(slice []int, num int) []int {
	newSlice := []int{}
	for _, v := range slice {
		if v != num {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
