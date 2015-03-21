// Bullcow	The Game of Bulls and Cows
//
//			Mar 2015
//	  		V 0.1
//
//		To Do:  seed Rand gen; check it generates a '9'
//				5679 does work properly ??

package main

import (
	"fmt"
	"os"
	"bufio"
	"math/rand"
	"strings"
)

//  The computer generated string of four digits is kept here
var target = make([]byte, 10)

// Latest move
var move = make([]byte, 10)

var gameOver bool
// In case the game runs in a loop, initialise (again)
func initVariables() {
	gameOver = false
}

//  Ask the player for his or her next move
func getMove() {
	for {
		m := askMove()
		if m == "QUIT" {
			os.Exit(0)
		}
		move = []byte(m)
		return
	}
}

// Ask the player for a guess (four digits)
//
func askMove() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your move (Four numbers 1 and 9, with no repeats) : ")
	line, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(1)		// Something seriously wrong
	}
	if len(line) != 5 {
		return "QUIT"		// Tell the caller that player wants out!
	}
	return line
}


// Make the initial four-digit string (no repeats)
func makeComputerTarget() string {
	m := byte(rand.Intn(8) + 1 + '0')
	s := string(m)
	for len(s) < 4 {
		m = byte(rand.Intn(8) + 1 + '0')
		if !strings.Contains(s, string(m)) {
			s += string(m)
		}
	}
	//fmt.Println(s)
	return s
}

func displayClues() {
	var cows, bulls int = 0, 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if target[i] == move[j] {
				if i == j {
					bulls++
				} else {
					cows++
				}
			}
		}
	}
	fmt.Printf("Bulls: %d     Cows: %d\n", bulls, cows)
	if bulls == 4 {
		gameOver = true
	}
	return	
}

//  Run the game

func main() {
	for {
		initVariables()
		target = []byte(makeComputerTarget())
		getMove() 	// Ask for valid move.
		displayClues()
		if gameOver == true {
			break
		}
	}
}
			

