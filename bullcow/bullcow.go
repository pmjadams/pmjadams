// Bullcow	The Game of Bulls and Cows
//
//			Mar 2015
//	  		V 0.1
//	  		V 0.2
//	  		V 0.3	add rand seed; convert byte -> string; screen user input
//			V 0.4	change rand seed to use 'Now' in nsecs
//			V 0.5	mask seed for rand. Make code Windows friendly
//
//			Note: rand.Intn(N) seems to generate numbers from 0 - (N-1), not 0 - N
//
//			ToDo: print out version number if asked

package main

import (
	"fmt"
	"os"
	"bufio"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//  The computer generated string of four digits is kept here
var target string

// Latest move by player
var move string

// On Windows the EOL is CR NL, otherwise it is NL
// The input line is read in up to and including the NL
// LENEOL is 1 for Windows, 0 for Unix-like, viz. the extra line length
var ΔEOL int = 0

//  Run the game

func main() {
	if runtime.GOOS == "windows" {
		ΔEOL = 1
	}
	initVariables()
	target = makeComputerTarget()
	for {
		getMove() 	// Ask for a valid move.
		displayClues()
		if gameOver == true {
			fmt.Println("Well done!")
			break
		}
	}
}
			
var gameOver bool

// In case the game runs in a loop, initialise (again)
func initVariables() {
	gameOver = false
	a := time.Now().UnixNano()
	b := a & 0xFFFFFFFF		// Mask
	rand.Seed(b)

}

//  Ask the player for his or her next move
func getMove() {
	for {
		if move = askMove(); move == "QUIT" {
			os.Exit(0)
		}
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
	ll := len(line) + ΔEOL
	if ll != 5 { 
		return "QUIT"		// Tell the caller that player wants out!
	}
	// Strip EOL
	l := line[0:4]
	// Use this fn to check a number has been input
	if _, err := strconv.Atoi(l); err != nil {
		fmt.Printf("Sorry! %s is not a number...", l)
		os.Exit(1)
	}
	// Now check there are no repeat digits
	// l is the four digits entered
	// m is a rotation of these digits
	var m string
	m = l
	for i:=0; i<3; i++ {
	  m = m[1:4] + m[0:1]
	  if (l[0:1]==m[0:1]) || (l[1:2]==m[1:2]) || (l[2:3]==m[2:3]) || (l[3:4]==m[3:4]) {
		fmt.Printf("Sorry! %s contains repeat digits...", l)
		os.Exit(1)
	  }
	}  
	return l
}

// Make the initial four-digit string (no repeats)
func makeComputerTarget() string {
	m := byte(rand.Intn(9) + 1 + '0')
	s := string(m)
	for len(s) < 4 {
		m = byte(rand.Intn(9) + 1 + '0')
		if !strings.Contains(s, string(m)) {
			s += string(m)
		}
	}
	return s
}

func displayClues() {
	var cows, bulls int = 0, 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if target[i:i+1] == move[j:j+1] {
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

