package echo

import (
	"flag"
	"fmt"
	"strings"
)

func Echo4() {
	n := flag.Bool("n", false, "omit trailing newline")
	sep := flag.String("s", " ", "separator")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
