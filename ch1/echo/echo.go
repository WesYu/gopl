package echo

import (
	"fmt"
	"os"
	"strings"
)

func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func Exercise_1_1() {
	fmt.Println(strings.Join(os.Args, " "))
}

func Exercise_1_2() {
	for i, arg := range os.Args {
		if i != 0 {
			fmt.Printf("Index: %v, argument: %q\n", i, arg)
		}
	}
}
