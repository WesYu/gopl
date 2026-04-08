package dup

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Count the number of times each string is read from the input file, ignoring
// potential errors occurred when reading from the input.
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

// Prints the text of each line that appears more than once in the standard
// input, preceded by its count. To finish, press ctrl+D to indicate end of file
// in the standard input.
func Dup1() {
	counts := make(map[string]int)
	countLines(os.Stdin, counts)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Prints the count and text of lines that appear more than once in the
// input. If there's no command line argument, it reads from stdin. Otherwise,
// it reads from a list of named files, with each argument as a file name.
func Dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Same as Dup2, except it only accepts files via the command line arguments as
// file names. Reads the entire input into memory in one big gulp, split it into
// lines all at once, then process the lines.
func Dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for line := range strings.SplitSeq(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Exercise 1.4: Modify Dup2 to print the names of all files in which each
// duplicated line occurs.
func DupName() {
	counts := make(map[string]int)
	filenames := make(map[string][]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for i, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "DupName: %v\n", err)
				continue
			}
			input := bufio.NewScanner(f)
			seen := make(map[string]bool)
			for input.Scan() {
				line := input.Text()
				counts[line]++
				if !seen[line] {
					filenames[line] = append(filenames[line], i)
					seen[line] = true
				}
			}
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)
			for _, i := range filenames[line] {
				fmt.Printf("%s ", os.Args[i+1])
			}
			fmt.Println()
		}
	}
}
