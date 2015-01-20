package main

import (
	"fmt"
	"os"
	"bufio"
)

var board = [][]byte{
	{'1', '2', '3'},
	{'4', '5', '6'},
	{'7', '8', '9'},
}
var play = [][]byte{
	{'.', '.', '.'},
	{'.', '.', '.'},
	{'.', '.', '.'},
}

func display(bd [][]byte) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Printf("%c ", bd[x][y])
		}
		fmt.Println("")
	}
}

var move byte

func interact() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your move (as a number between 1 and 9) : ")
	line, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	if _, err := fmt.Sscanf(line, "%c", &move); err != nil {
		fmt.Fprintln(os.Stderr, "invalid input")
	}
	fmt.Printf("\nThe move was: %c ", move)
	fmt.Println("")
}


func main() {

	display(board)
	fmt.Println("")
	display(play)
	fmt.Println("")
	interact()
}
