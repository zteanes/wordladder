package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func testArgs(args []string) {
	// ensure correct amount of args was given
	if len(args) != 3 {
		fmt.Println("\nUsage: wordladder <dictionary> <start> <end>")
		os.Exit(1)
	}

	// ensure the start and end words are the same length
	if len(args[1]) != len(args[2]) {
		fmt.Println("\nThe start and end words must be the same length.")
		os.Exit(1)
	}

	// ensure the dictionary files exists
	//
	if _, err := os.Stat(args[0]); errors.Is(err, fs.ErrNotExist) {
		fmt.Println("\nThe dictionary file does not exist.")
		os.Exit(1)
	}
}

// This is our main function for entering the program. It will take command line
// arguments and then call the appropriate functions to build the word ladder.
func main() {
	// get the arguments and see if they are valid
	args := os.Args[1:]
	testArgs(args)

	fmt.Println("Hello, World!")
}
