//	short program to see what input the keyboard provides
//	this version does not use raw mode (blocks for NL)
//	raw input version
//	V.2	6 Feb 2015	PA

package main

import (
        "fmt"
        "os"
        "bufio"
)

import "https://golang.org/x/crypto/ssh/terminal"

// Ask the user to press a key
//
func askKey() byte {
	var ch byte
	// put stdin into raw mode
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(0, oldState)
	// now read a byte
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Please press a key: ")
    if chr, err := reader.ReadByte();  err != nil {
        os.Exit(1)              // Something seriously wrong
    } else {
		ch = chr
	}
	return ch
}


// Main test loop
func main() {
	c := askKey()
	fmt.Printf("The char returned was %o\n", c)
}

