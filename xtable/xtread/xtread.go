//	xtread:  Process xt games records
//				Mar 2015,  Ver. 0.2
//
//				Ver. 0.1	-	implement the basics, plus debug output
//				Ver. 0.2	-	prints start time and duration
//
//				ToDo: work out what stats to produce
//						

package main

import (
		"os"
		"fmt"
		"log"
		"bufio"
		"encoding/gob"
)

// One Game consists of a number of x times y 'plays'.
// A player may play a number of Games.

type Game struct {
	Name	string
	Id		int
	NGoes	int
	Start	int64
	End		int64
	Plays	[]*Play
}

// The Play is one 'x time y' question.
// The answer given by the player is the Guess.
// Result records if the Guess was correct.

type Play struct {
	Id		string
	A		int
	B		int
	Guess	int
	Result	bool
}

// A custom method to override the default.
func (play Play) String() string {
	return fmt.Sprintf("\nId: %s, n1: %d, n2: %d, Guess: %d, Result: %t", play.Id, play.A, play.B, play.Guess, play.Result)
}

// Not used, yet
const (
		magicNumber = 0x125E
		fileVersion = 100
)

// utility function
func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func main () {
  // Read from stdin
	r1 := bufio.NewReader(os.Stdin)
	loop := true
	for loop {	// per game loop
		games := Game{}
  		e := readGame(r1, &games)
  		if e != nil {	// EOF reached
  			loop = false
  		} else {
  			games.End -= games.Start
			fmt.Println(games)
		}
	}
	return
}

func readGame(rd *bufio.Reader, gm *Game)  error {
	decoder := gob.NewDecoder(rd)
	err := decoder.Decode(gm)
	return err
}
