package main

import (
	"os"
	"bufio"
	"fmt"
)

func main()  {
	files := os.Args[1:]
	if len(files) == 0{
		os.Exit(0)
	} else {
		for _, file := range files {
			counts := make(map[string] int)
			f, err := os.Open(file)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				os.Exit(1)
			}
			countLines(f, counts)
			flag := false
			for _, count := range counts {
				if count > 1 {
					flag = true
				}
			}
			if flag {
				fmt.Printf("%s has duplicated rows\n", file)
			}
		}
	}
}

func countLines(f *os.File, counts map[string] int)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
