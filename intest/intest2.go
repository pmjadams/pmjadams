package main

// This doesn't work either.
// 6 Feb 2015

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    // disable input buffering
    exec.Command("stty", "-F", "/dev/ttys000", "cbreak", "min", "1").Run()
    // do not display entered characters on the screen
    exec.Command("stty", "-F", "/dev/ttys000", "-echo").Run()
    // restore the echoing state when exiting
    defer exec.Command("stty", "-F", "/dev/ttys000", "echo").Run()

    var b []byte = make([]byte, 1)
    for {
        os.Stdin.Read(b)
        fmt.Println("I got the byte", b, "("+string(b)+")")
    }
}

// messy