package main

import (
	"time"
	"os"
	"fmt"
	"strings"
)

func main()  {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("%s Echo finished at: %.10f\n", s, time.Since(start).Seconds())

	start = time.Now()
	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("%s Echo finished at: %.10f\n", s, time.Since(start).Seconds())

	start = time.Now()
	fmt.Printf("%s Echo finished at: %.10f\n",
		strings.Join(os.Args[1:], " "), time.Since(start).Seconds())
}
