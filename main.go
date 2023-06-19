package main

import "fmt"

var board [3][3]string
var currentPlayer string = "X"

func main() {
	fmt.Println("Welcome to Tic-Tac-Toe!")
	initializeBoard()
	playGame()
}

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "-"
		}
	}
}

func makeMove(row, col int) {
	if board[row][col] == "-" {
		board[row][col] = currentPlayer
	} else {
		fmt.Println("Invalid move. Try again")
	}
}

func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}

func displayBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%s ", board[i][j])
		}
		fmt.Println()
	}
}

func checkGameStatus() string {
	// Dikey ve yatay kombinasyonları kontrol et
	for i := 0; i < 3; i++ {
		if board[i][0] != "-" && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			return board[i][0]
		}
		if board[0][i] != "-" && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			return board[0][i]
		}
	}

	// Çapraz kombinasyonları kontrol et
	if board[0][0] != "-" && board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] != "-" && board[0][2] == board[1][1] && board[0][2] == board[2][0] {
		return board[0][2]
	}

	// Beraberlik durumunu kontrol et
	isFull := true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "-" {
				isFull = false
				break
			}
		}
		if !isFull {
			break
		}
	}
	if isFull {
		return "draw"
	}

	// Oyun devam ediyor
	return "ongoing"
}

func playGame() {
	for {
		fmt.Printf("Current player: %s\n", currentPlayer)
		fmt.Print("Enter row index (0-2): ")
		var row int
		fmt.Scanln(&row)

		fmt.Print("Enter column index (0-2): ")
		var col int
		fmt.Scanln(&col)

		makeMove(row, col)
		switchPlayer()

		fmt.Println("Updated Board:")
		displayBoard()

		status := checkGameStatus()
		if status == "X" || status == "O" {
			fmt.Printf("Player %s wins!\n", status)
			break
		} else if status == "draw" {
			fmt.Println("The game is a draw.")
			break
		}
	}
}

//test
