package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

func main()  {
	counts := make(map[string] int)
	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, count := range counts {
		fmt.Printf("%d\t%s\n", count, line)
	}
}
