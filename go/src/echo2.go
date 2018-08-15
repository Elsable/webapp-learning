package main

import (
	"os"
	"fmt"
)

func main()  {
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx, " ", arg)
	}
}
