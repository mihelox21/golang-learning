package main

import (
	"fmt"
	"strings"
)

type GameBoard = [][]string
type BoardRow = []string
type GameSymbol = string
type Patterns = [][]int

type PointPosition struct {
	row    uint
	column uint
}

var winPatterns Patterns = Patterns{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

var board GameBoard = GameBoard{
	BoardRow{"_", "_", "_"},
	BoardRow{"_", "_", "_"},
	BoardRow{"_", "_", "_"},
}

var currentSymbol string = "X"

func transformTo2dPosition(position int) (int, int) {
	row := position / 3
	column := position - 3*row

	return row, column
}

func isGameBoardEmpty() bool {
	spaces := 0
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			currentField := board[row][col]
			if currentField == "X" || currentField == "O" {
				spaces += 1
			}
		}
	}
	return spaces == 9
}

func playerWonGame() bool {
	for line := 0; line < len(winPatterns); line++ {
		matchedItems := 0

		for item := 0; item < len(winPatterns[line]); item++ {
			row, col := transformTo2dPosition(winPatterns[line][item])

			if board[row][col] == currentSymbol {
				matchedItems += 1
			}

			if matchedItems == 3 {
				return true
			}
		}
	}
	return false
}

func nextGameSymbol(previousSymbol string) string {
	if previousSymbol == "X" {
		return "O"
	} else {
		return "X"
	}
}

func updateSymbol(newValue string) {
	currentSymbol = newValue
}

func updateBoardField(board GameBoard, position PointPosition, symbol string) {
	board[position.row][position.column] = symbol
}

func readUserMove() PointPosition {
	var userChoice PointPosition
	_, err := fmt.Scan(&userChoice.row, &userChoice.column)
	if err != nil || userChoice.row >= 3 || userChoice.column >= 3 {
		clearConsole()
		fmt.Println("Wrong position.. try again..")
		return readUserMove()
	}
	return userChoice
}

func drawBoard(board GameBoard) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func clearConsole() {
	fmt.Println("\033[2J")
}

func gameEngine(board GameBoard, currentSymbol string) {
	clearConsole()
	drawBoard(board)
	fmt.Println("Enter coordinates separated by space | x = row, y = column")
	userMove := readUserMove()
	updateBoardField(board, userMove, currentSymbol)
}

func main() {
	for {
		gameEngine(board, currentSymbol)
		if isGameBoardEmpty() {
			fmt.Println("Draw!")
			break
		}
		if playerWonGame() {
			fmt.Printf("Player with symbol %s won the game!", currentSymbol)
			break
		}
		updateSymbol(nextGameSymbol(currentSymbol))
	}
	fmt.Println("Game finished.")
}
