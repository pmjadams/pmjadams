// The Game of Noughts and Crosses
//
//	Jan 2015
//	  V 0.75

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


var win = []byte{   // There are 8 rows of possible winning lines.
	1, 2, 3,		// To use these numbers as indicies, subtract 1.
	4, 5, 6,
	7, 8, 9,
	1, 4, 7,
	2, 5, 8,
	3, 6, 9,
	1, 5, 9,
	3, 5, 7,
}


//  If a winning row contains 3 'x's, adding the 3 as numbers equals 360.
//  (The ascii letter small x has decimal value of 120)
//  Similarly, the decimal value of 3 'o's is 333
//  No other combination of 'x's, 'o's and 0's, adds up to these two nos.

var gameOver bool

var move byte

const QUIT = 55

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
	for {
		m := askMove()
		if m == QUIT {
			os.Exit(0)
		}
		if m < 0 || m > 8 {	// indices, not play squares
	fmt.Println("Invalid input. Please enter a number between 1 and 9, and press return")
	 	// <need some logic here>
		continue
		}	
		// is this move allowed? Not if square is occupied.
		if play[m] != 0 {
			fmt.Println("That square is occupied. Please try again.")
		} else {
			move = m
			return
		}
	}
}

// Ask the player to play a move, a number between 1 and 9
//
func askMove() byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your move (as a number between 1 and 9) : ")
	line, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(1)		// Something seriously wrong
	}
	if len(line) != 2 {
		return QUIT		// Tell the caller that player wants out!
	}
	mv := line[0]
	
	return mv - '1'		// return the index into play[], corresponding to desired move
}


func updateAndDisplayBoard(mv byte) {
	play[move] = mv
	displayBoard()
}

func moveWinsGame() bool {
	sum := 0
	for i := 0; i < 8; i++ {
		sum = 0
		for j := 0; j < 3; j++ {
			sum += int(play[win[3*i + j] - 1])
		}
		// fmt.Printf("sum is %d: \n", sum)
		if sum == 360 {		// 'x' wins
			fmt.Println(" 'x' wins!")
			return true
		} else if sum == 333 {	// 'o' wins!
			fmt.Println(" 'o' wins!")
			return true
		} 
		
	}
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
			

