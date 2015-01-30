// The Game of Noughts and Crosses
//
//	Jan 2015
//	  V 0.7

package main

import (
	"fmt"
	"os"
	"bufio"
	"math/rand"
)

//  The squares are numbered like this

var board = []byte{
	'1', '2', '3',
	'4', '5', '6',
	'7', '8', '9',
}

//  The moves are recorded here:  'x', 'o' or 0 for empty

var play = make([]byte, 9)

var gameOver bool

var move byte

func initVariables() {

	gameOver = false
	for item := range play {
		play[item] = 0
	}
}

//  Display the board, or the moves, on the screen

func displayBoard() {
	display(board)
	fmt.Println("")
	display(play)
	fmt.Println("")
}

func display(bd []byte) {
	for item := range bd {
		printBox(bd[item])
		switch item {
			case 2, 5, 8: fmt.Println("")
		}
	}
}

func printBox(sq byte) {
	if sq == 0 {
		fmt.Printf(". ")
	} else {
		fmt.Printf("%c ", sq)
	}
}



//  Ask the player for his or her next move

func getMove() {
	reader := bufio.NewReader(os.Stdin)
AGAIN:	fmt.Printf("Enter your move (as a number between 1 and 9) : ")
	line, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}
	if _, err := fmt.Sscanf(line, "%c", &move); err != nil {
		fmt.Fprintln(os.Stderr, "invalid input\n")
		os.Exit(1)
	}
	fmt.Printf("value of move is %d\n", move)
	if move > 127 {
		os.Exit(0)
	}
	if move <= '0' || move > '9' {
	 fmt.Println("Invalid input. Please enter a number between 1 and 9, and press return")
	}
	fmt.Printf("\nThe move was: %d ", move - '0')
	fmt.Println("")
	move = move - '0' - 1
	// is this move allowed? Not if square is occupied.
	if play[move] != 0 {
		fmt.Println("That move is not allowed. Please try again.")
	} else {
		return
	}
	goto AGAIN
}

func updateAndDisplayBoard(mv byte) {
	play[move] = mv
	displayBoard()
}

func moveWinsGame() bool {

	return false
	
}

func makeComputersMove() {
	m := byte(rand.Intn(8))
	for  play[m] != 0 {
		 m = byte(rand.Intn(8))		
	}
	move = m
	return
}

func playAnotherGame() bool {
	reader := bufio.NewReader(os.Stdin)
AGIN: fmt.Printf("Play another game? (y/n) : ")
	line, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}
	if _, err := fmt.Sscanf(line, "%c", &move); err != nil {
		fmt.Fprintln(os.Stderr, "invalid input")
	}
	if move == 'y' {
		return true
	} else if move == 'n' {
		return false
	} else {
		 fmt.Println("Please enter y or n, and press return")
		fmt.Printf("\nDebug: The move was: %c ", move)
		fmt.Println("")
		goto AGIN
	}
}



//  Run the game

func main() {

	for {
		initVariables()
		displayBoard()
		for gameOver == false {
			getMove() 	// Ask for valid move.
			updateAndDisplayBoard('o')
			if moveWinsGame() == true {
				gameOver = true
				break
			}
			makeComputersMove()
			updateAndDisplayBoard('x')
			if moveWinsGame() == true {
				gameOver = true
				break
			}
		}
		if playAnotherGame() == false {
			break
		}
	}
	return
}
			

