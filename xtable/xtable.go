//	xtable:  Practise your times table
//				Mar 2015,  Ver. 0.8
//
//				Ver. 0.1	-	implement one play of the game
//				Ver. 0.2	-	use Cursor codes to print success (or not) on
//								same line as answer.
//				Ver. 0.3	-	trying to create structs to hold play data
//				Ver. 0.4	-	tidy up, fill in Game struct.
//				Ver. 0.5	-	still tidying up
//				Ver. 0.6	-	starting to data files items, debugging
//				Ver. 0.7	-	can create/open file and append write to it
//				Ver. 0.8	-	archive games to disk using gob encoding
//
//				ToDo: 
//						Write prog to read in games and display stats
//						Then tidy up the code

package main

import (
		"os"
		"fmt"
		"log"
		"bufio"
		"math/rand"
		"time"
		"strconv"
		"strings"
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
	return fmt.Sprintf("\n<Id> %s, <n1> %d, <n2> %d, <Guess> %d, <Result> %t", play.Id, play.A, play.B, play.Guess, play.Result)
}

var	n1, n2 int
var ans int

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
  //construct the filename for datafile
	var t time.Time
	var ss, filename string
	var fields []string
	t = time.Now()
	ss = t.Format(time.RFC3339)
	fields = strings.Split(ss, "T")
	filename = "xt" + fields[0] + ".xtd"
  // Open this file, create it if not exists
	fmt.Println(filename)
	f, err := os.Open(filename)
    if err != nil && os.IsNotExist(err) {
    	f, err = os.Create(filename)
    	check(err)
    }
    err = f.Chmod(0666)
    check(err)
  // close the now existing file
  	err = f.Close()
  	check(err)
  // Open it for writing
  	f, err = os.OpenFile(filename, 01, 0666)
  	check(err)
  // Seek to the end
  	_, error := f.Seek(int64(0), 2)
  	check(error)
  // File ready for writing next blob
  
  // Not quite
  	r1 := bufio.NewWriter(f)
  /*	_, err = fmt.Fprintf(r1, "Hello File \n")
  	check(err)
  	err = r1.Flush()
 	check(err)  	*/
  // Ready for the game
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
			reply := dialog("Another game? (Y/n): ")
			reply = strings.TrimSpace(reply)
			wantsToPlay = (reply != "n")
			games.End = time.Now().Unix()
			// print out value of games and plays
			fmt.Println(games)
			fmt.Printf("The game took %d seconds\n", games.End - games.Start)
			// write game details to g
			writeGame(r1, &games)
			err = r1.Flush()
			check(err)
		} // end of for wantsToPlay
		reply := dialog("A new player? (Y/n): ")
		reply = strings.TrimSpace(reply)
		newGame = (reply != "n")
	}
  	err = f.Close()
  	check(err)

}

func writeGame(r1 *bufio.Writer, gm *Game) {
	encoder := gob.NewEncoder(r1)
	encoder.Encode(gm)
	return
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
	id := dialog("  Id number: ")
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
	var num int
	var err error
	goes := dialog("How many goes? : ")
	goes = strings.TrimSpace(goes)
	num, err = strconv.Atoi(goes); 
	if err != nil {
		fmt.Println("Sorry, not a number!!")
		os.Exit(1)		// for now
	}
	return num
}