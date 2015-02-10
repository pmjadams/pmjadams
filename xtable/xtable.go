//	xtable:  Practise your times table
//				Feb 2015,  Ver. 0.2
//
//				Ver. 0.1	-	implement one play of the game
//				Ver. 0.2	-	use Cursor codes to print success (or not) on
//								same line as answer.
//
//				ToDo: create structs to hold play and game data,
//					then add archiving of games to disk.
//					Develop the game to record players, offer a number of
//					plays per game, e.g. 30, and play against the clock,
//					e.g. 60 seconds.

package main

import (
		"os"
		"fmt"
		"bufio"
		"math/rand"
		"time"
)

var n1, n2 int


func main () {
	var ans int					// holds the player's answer
	s := "\033[A\033[20C"		
	t := time.Now()
	rand.Seed(t.Unix())
	//  init the first two numbers
	n1 = rand.Intn(10) + 2		// a number in the range 2...12
	n2 = rand.Intn(10) + 2		// again
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("   %d x %d = ", n1, n2)		// make the challenge
	s, err := reader.ReadString('\n')
	if  err != nil {		// read in the answer
		os.Exit(1)											// for the moment
	}
	if _, err := fmt.Sscanf(s, "%d", &ans); err != nil {
		fmt.Fprintln(os.Stderr, "invalid input\n")
		os.Exit(1)
	}
	// test for the correct answer
	if n1*n2 != ans {
		fmt.Println("\033[A\033[20CIncorrect")		// Cursor Up and Forward 20 places
	} else {
		fmt.Println("\033[A\033[20CCorrect")		// Ditto
	}
}
