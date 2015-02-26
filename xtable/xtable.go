//	xtable:  Practise your times table
//				Feb 2015,  Ver. 0.3
//
//				Ver. 0.1	-	implement one play of the game
//				Ver. 0.2	-	use Cursor codes to print success (or not) on
//								same line as answer.
//				Ver. 0.3	-	trying to create structs to hold play data
//				Ver. 0.4	-	tidy up, fill in Game struct.
//				Ver. 0.5	-	still tidying up
//
//				ToDo: 
//					Add archiving of games to disk.
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
		"strconv"
		"strings"
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
	return fmt.Sprintf("\n<Id> %s, <n1> %d, <n2> %d, <Guess> %d, <Result> %t", play.Id, play.A, play.B, play.Guess, play.Result)
}

var	n1, n2 int
var ans int


func main () {
	newGame := true
	for newGame {
		games := Game{}
	// Ask for credentials
		whoAreYou(&games)
	// Instruct the player
		instruct()
	// Main Loop control
		wantsToPlay := true
		for wantsToPlay {
			count := askNoOfSums()
			games.NGoes = count
	// Make storage for this number
			plays := make([]*Play, count)
			games.Plays = plays
			games.Start = time.Now().Unix()
	// Seed Random
			t := time.Now()
			rand.Seed(t.Unix())
	// Play loop
			for i := 0; i < count; i++ {
			//  init the first two numbers
				n1 = rand.Intn(10) + 2		// a number in the range 2...12
				n2 = rand.Intn(10) + 2		// again
			// init the play
				plays[i] = new(Play)
			//	Make the challenge
				fmt.Printf("   %d x %d = ", n1, n2)
				plays[i].Id = strconv.Itoa(games.Id)
				plays[i].A = n1
				plays[i].B = n2
				
			// Get the reply
				s := dialog("")
				if _, err := fmt.Sscanf(s, "%d", &ans); err != nil {
					fmt.Fprintln(os.Stderr, "invalid input\n")
					os.Exit(1)
				}
				games.Plays[i].Guess = ans
				if n1*n2 != ans {
				// Cursor Up and Forward 20 places
					fmt.Println("\033[A\033[20CIncorrect")	
					plays[i].Result = false
				} else {
					fmt.Println("\033[A\033[20CCorrect")		// Ditto
					games.Plays[i].Result = true
				}
			} // end of for i
			wantsToPlay = false
			games.End = time.Now().Unix()
			// print out value of games and plays
			fmt.Println(games)
			fmt.Printf("The game took %d seconds\n", games.End - games.Start)
		} // end of for wantsToPlay
	}
}

func dialog(str string) string {
	fmt.Printf(str)
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if  err != nil {		// read in the answer
		os.Exit(1)			// for the moment
	}
	return s
}

func instruct() {
}

func whoAreYou(g *Game) {
	var num int
	var err error
	nm := dialog("To play this game, please enter your name and Id number\n  Name: ")
	id := dialog("Id number: ")
	nm, id = strings.TrimSpace(nm), strings.TrimSpace(id)
	if len(nm) == 0 {
		fmt.Println("No name entered!!")
		os.Exit(1)		// for now
	}
	g.Name = nm
	num, err = strconv.Atoi(id); 
	if err != nil {
		fmt.Println("Sorry, not a number!!")
		os.Exit(1)		// for now
	}
	g.Id = num
}

func askNoOfSums() int {
	return 6
}