//	recordcount:  Read book data records and count them
//				Mar 2015,  Ver. 0.1
//
//				Ver. 0.1	-	implement the basics, plus debug output
//				
//

package main

import (
	"bufio"
	"fmt"
	"os"
)


// utility function
func check(e error) {
	if e != nil {
		panic(e)
	}
}


var j int //  Indexes through lb records


func main() {
	// Read
	f, err := os.Open("../books.sql")
	check(err)

	// Set input source.
	scanner := bufio.NewScanner(f)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanLines)
	parsing := true
	j = 0
	// Loop begins here
	for parsing {
		// Advance to next line, & validate the input
		if scanner.Scan() == false {
			break // EOF, hopefully
		}
		s := scanner.Text()
		if s == "" || s[0] != '(' {
			continue // Not a data record
		}
		j++
		if j == 5999 {
			parsing = false
		}
	}
	fmt.Printf("There are %d records.\n", j)
}

