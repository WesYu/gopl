package dup

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// dup1 prints the text of each line that appears more than once in the standard
// input, preceded by its count.
func Dup1() {
	counts := make(map[string]int)
	countLines(os.Stdin, counts)
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// dup2 prints the count and text of lines that appear more than once in the
// input. It reads from stdin or from a list of named files.
func Dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
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

// Modify dup2 to print the names of all files in which each duplicated line
// occurs.
func Exercise_1_4() {
	counts := make(map[string]int)
	filenames := make(map[string][]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for i, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
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

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

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
