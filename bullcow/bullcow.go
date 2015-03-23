// Bullcow	The Game of Bulls and Cows
//
//			Mar 2015
//	  		V 0.1
//	  		V 0.2
//	  		V 0.3	add rand seed; convert byte -> string; screen user input
//
//			Note: rand.Intn(N) seems to generate nos from 0 - (N-1), not N
//				

package main

import (
	"fmt"
	"os"
	"bufio"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//  The computer generated string of four digits is kept here
var target string

// Latest move by player
var move string

var gameOver bool
// In case the game runs in a loop, initialise (again)
func initVariables() {
	gameOver = false
	rand.Seed(time.Now().Unix())
}

//  Ask the player for his or her next move
func getMove() {
	for {
		m := askMove()
		if m == "QUIT" {
			os.Exit(0)
		}
		move = m
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
	// Strip NL
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

//  Run the game

func main() {
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
			

