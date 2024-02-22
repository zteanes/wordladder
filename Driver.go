package wordladder

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func testArgs(args []string) bool {
	goodArgs := true
	// ensure correct amount of args was given
	if len(args) != 3 {
		fmt.Println("Usage: wordladder <dictionary> <start> <end>")
		goodArgs = false
	}

	// ensure the start and end words are the same length
	if len(args[1]) != len(args[2]) && goodArgs {
		fmt.Println("The start and end words must be the same length.")
		goodArgs = false
	}

	// ensure the dictionary files exists
	//
	if _, err := os.Stat(args[0]); errors.Is(err, fs.ErrNotExist) && goodArgs {
		fmt.Println("The dictionary file does not exist.")
		goodArgs = false
	}

	return goodArgs
}

func main() {
	// get the arguments and see if they are valid
	args := os.Args[1:]
	if testArgs(args) {

	}

	fmt.Println("Hello, World!")
}
