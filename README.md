# sdku

# Sudoku Solver in Go

## Overview
This repository contains a Go implementation of a Sudoku solver. It utilizes a backtracking algorithm enhanced with constraint propagation for efficient solving of Sudoku puzzles.

## Files
- `main.go`: The main file that contains the logic for parsing input, invoking the solver, and displaying the output.
- `updatepossibilities.go`: Part of the `piscine` package, this file includes the `UpdatePossibilities` function that implements constraint propagation to narrow down potential values for each cell in the Sudoku grid, in order to optimize solution.

## Setup
To set up and run the Sudoku solver, follow these steps:

1. **Clone the Repository:**
   ```
   <!-- git clone https://zero.academie.one/git/adunayev/sudoku -->
   cd sudoku
   ```
2. **Run the Solver:**
   - You can run the solver by passing a Sudoku puzzle as command-line arguments. For example:
   
   ```
   cd main/
   go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
   ```

## Usage
- Sudoku puzzles should be input as 9 strings of length 9, representing each row of the puzzle.
- Use a period (`.`) to represent empty cells.
- The program will output the solved puzzle or an error message if the puzzle is invalid or unsolvable.

## Contributing
This is an exercise to submit in the course of our raid. There are three contributors and authors to this particular project: adunayev, aaidynul, aabdrama