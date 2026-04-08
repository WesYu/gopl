package echo

import (
	"fmt"
	"os"
	"strings"
)

// Prints the command-line arguments.
func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// Prints the command-line arguments.
func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// Prints the command-line arguments.
func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// Exercise 1.1: Modify the echo program to also print os.Args[0], the name of
// the command that invoked it.
func EchoWithCommandName() {
	fmt.Println(strings.Join(os.Args, " "))
}

// Exercise 1.2: Modify the echo program to print the index and value of each of
// its arguments, one per line.
func EchoWithIndex() {
	for i, arg := range os.Args {
		if i != 0 {
			fmt.Printf("Index: %v, argument: %q\n", i, arg)
		}
	}
}
