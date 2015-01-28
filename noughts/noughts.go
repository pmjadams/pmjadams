// The Game of Noughts and Crosses
//
//	Jan 2015
//	  V 0.5

package main

import (
	"fmt"
	"os"
	"bufio"
)

//  The squares are numbered like this

var board = [][]byte{
	{'1', '2', '3'},
	{'4', '5', '6'},
	{'7', '8', '9'},
}

//  The moves are recorded here:  X,  0 or '.'

var play = [][]byte{
	{'.', '.', '.'},
	{'.', '.', '.'},
	{'.', '.', '.'},
}

//  Display the board, or the moves, on the screen

func display(bd [][]byte) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Printf("%c ", bd[x][y])
		}
		fmt.Println("")
	}
}

var move byte

//  Ask the player for his or her next move

func interact() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your move (as a number between 1 and 9) : ")
	line, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	if _, err := fmt.Sscanf(line, "%c", &move); err != nil {
		fmt.Fprintln(os.Stderr, "invalid input")
	}
	if move <= '0' || move > '9' {
	 fmt.Println("Invalid input. Please enter a number between 1 and 9, and press return")
	}
	fmt.Printf("\nThe move was: %d ", move - '0')
	fmt.Println("")
}


//  Run the game

func main() {

	display(board)
	fmt.Println("")
	display(play)
	fmt.Println("")
	interact()
}
