//	booklib:  Process book records in sql format, in CSV format.
//				Mar 2015,  Ver. 0.1
//
//				Ver. 0.1	-	implement the basics, plus debug output
//				Ver. 0.2	-	read in one record of data
//				Ver. 0.3	-	modify SplitN, seems to work :-)
//				Ver. 0.4	-	simplify layout, remove reps
//
//				ToDo: Marshall records to disk
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Book Record
type BookRec struct {
	Book_Id          int
	Author           string
	Title            string
	Publisher        string
	Pub_Place        string
	Pub_Year         int
	Cn_Type          string
	Cn_Source        string
	Cn_Item          int
	Cn_Suffix        int
	ISBN             string
	ISSN             string
	URL              string
	Pages            string
	Copies           int
	Comment          string
	Comment_Date     int64
	Barcode          int
	HomeLib          string
	HoldingLib       string
	Onloan           int64
	DateLastBorrowed int64
	NotForLoan       bool
	Damaged          bool
	ItemLost         bool
	Withdrawn        bool
	Restricted       bool
	ItemNotes        string
	Issues           int
	Renewals         int
	Reserves         int
	CopyNumber       int
	Created_At       int64
}

// utility function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

var lb = make([]BookRec, 5000)
var j int //  Indexes through lb records
var record BookRec

func main() {
	// Read
	f, err := os.Open("books.sql")
	check(err)
	// Define some counters and other temp vars
	var i int    //	Indexes through fields in BookRec
	var s string // Temp place to hold a string
	var k int    // Holds result of Atoi()
	// Set input source.
	scanner := bufio.NewScanner(f)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanLines)
	// Make a slice to hold all the book records
	// Holder for a split string
	var bits []string
	// A Time holder
	var tim time.Time
	//
	parsing := true
	j = 0
	// Loop begins here
	for parsing {
		// Create a record
		record = BookRec{}
		// Advance to next line, & validate the input
		if scanner.Scan() == false {
			break // EOF, hopefully
		}
		s = scanner.Text()
		if s[0] != '(' {
			continue // Not a data record
		}
		s = s[1 : len(s)-2] // remove enclosing ( )

		bits = genSplit(s, ",", "'", 0, 33)
		if err = scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s", err)
		}
		// Trim for spaces
		for index := range bits {
			bits[index] = strings.TrimSpace(bits[index])
		}
		// Try parsing
		i = 0
		// Book_Id
		record.Book_Id, err = strconv.Atoi(bits[i])
		check(err)
		i++
		// Author
		record.Author = strings.Trim(s, "'")
		i++
		// Title
		record.Title = strings.Trim(s, "'")
		i++
		// Publisher
		record.Publisher = strings.Trim(s, "'")
		i++
		// Pub_Place
		record.Pub_Place = strings.Trim(s, "'")
		i++
		// Pub_Year
		record.Pub_Year, err = strconv.Atoi(bits[i])
		check(err)
		i++
		// Cn_Type
		record.Cn_Type = strings.Trim(bits[i], "'")
		i++
		// Cn_Source
		record.Cn_Source = strings.Trim(bits[i], "'")
		i++
		// Cn_Item
		record.Cn_Item, err = strconv.Atoi(bits[i])
		check(err)
		i++
		// Cn_Suffix
		record.Cn_Suffix, err = strconv.Atoi(bits[i])
		check(err)
		i++
		// ISBN
		record.ISBN = strings.Trim(bits[i], "'")
		i++
		// ISSN
		record.ISSN = strings.Trim(bits[i], "'")
		i++
		// URL
		record.URL = strings.Trim(bits[i], "'")
		i++
		// Pages (as a string)
		record.Pages = strings.Trim(bits[i], "'")
		i++
		// Copies
		record.Copies, err = strconv.Atoi(bits[i])
		check(err)
		i++
		// Comment
		record.Comment = strings.Trim(bits[i], "'")
		i++
		// Comment_Date
		record.Comment_Date = 0
		i++ //Ignore data input
		// Barcode
		record.Barcode, err = strconv.Atoi(bits[i])
		check(err)
		i++
		// HomeLib
		record.HomeLib = strings.Trim(bits[i], "'")
		i++
		// HoldingLib
		record.HoldingLib = strings.Trim(bits[i], "'")
		i++
		// Onloan
		s = strings.Trim(bits[i], "'")
		if s == "NULL" {
			record.Onloan = 0
		} else {
			tim, err = time.Parse("2006-01-02", s)
			check(err)
			record.Onloan = tim.Unix()
		}
		i++
		// DateLastBorrowed
		s = strings.Trim(bits[i], "'")
		if s == "NULL" {
			record.DateLastBorrowed = 0
		} else {
			tim, err = time.Parse("2006-01-02", s)
			check(err)
			record.DateLastBorrowed = tim.Unix()
		}
		i++
		// NotForLoan
		k, err = strconv.Atoi(bits[i])
		record.NotForLoan = (k == 1)
		check(err)
		i++
		// Damaged
		k, err = strconv.Atoi(bits[i])
		record.Damaged = (k == 1)
		check(err)
		i++
		// ItemLost
		k, err = strconv.Atoi(bits[i])
		record.ItemLost = (k == 1)
		check(err)
		i++
		// Withdrawn
		k, err = strconv.Atoi(bits[i])
		record.Withdrawn = (k == 1)
		check(err)
		i++
		// Restricted
		k, err = strconv.Atoi(bits[i])
		record.Restricted = (k == 1)
		check(err)
		i++
		// ItemNotes
		record.ItemNotes = strings.Trim(bits[i], "'")
		i++
		// Issues
		s = bits[i]
		if s == "NULL" {
			record.Issues = 0
		} else {
			k, err = strconv.Atoi(s)
			check(err)
			record.Issues = k
		}
		i++
		// Renewals
		s = bits[i]
		if s == "NULL" {
			record.Renewals = 0
		} else {
			k, err = strconv.Atoi(s)
			check(err)
			record.Renewals = k
		}
		i++
		// Reserves
		s = bits[i]
		if s == "NULL" {
			record.Reserves = 0
		} else {
			k, err = strconv.Atoi(s)
			check(err)
			record.Reserves = k
		}
		i++
		// CopyNumber
		record.CopyNumber, err = strconv.Atoi(bits[i])
		check(err)
		i++
		// Created_At
		s = strings.Trim(bits[i], "'")
		tim, err = time.Parse("2006-01-02 15:04:05", s)
		check(err)
		record.Created_At = tim.Unix()
		// We are finished
		lb[j] = record
		j++
		if j == 4999 {
			parsing = false
		}
	}
}

// Generic split (from strings std lib), modified: splits after each instance of sep,
// including sepSave bytes of sep in the subarrays. Does not
// look for sep inside quoted strings, quote is the char that
// brackets a string. Quote char is not stripped.
// The count determines the number of substrings to return:
//   n > 0: at most n substrings; the last substring will be the unsplit remainder.
//   n == 0: the result is nil (zero substrings)
func genSplit(s, sep, quote string, sepSave, n int) []string {
	if n <= 0 {
		return nil
	}
	if sep == "" || sep == quote {
		return nil
	}
	//if n < 0 {
	//	n = Count(s, sep) + 1
	//}
	toggle := false // Toggle when quote char seen
	c := sep[0]
	q := quote[0]
	start := 0
	a := make([]string, n)
	na := 0
	for i := 0; i+len(sep) <= len(s) && na+1 < n; i++ {
		if s[i] == q && (len(quote) == 1 || s[i:i+len(quote)] == quote) {
			toggle = (toggle == false)
		}
		if s[i] == c && !toggle && (len(sep) == 1 || s[i:i+len(sep)] == sep) {
			a[na] = s[start : i+sepSave]
			na++
			start = i + len(sep)
			i += len(sep) - 1
		}
	}
	a[na] = s[start:]
	return a[0 : na+1]
}
