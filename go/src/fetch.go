package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
	"strings"
)

func main() {
	for _, file := range os.Args[1:] {
		if strings.HasPrefix(file, "http://") {

		} else if strings.HasPrefix(file, "https://") {

		} else {
			file = "http://" + file
		}
		fetchCopy(file, os.Stdout)
	}
}

func fetchCopy(url string, out io.Writer)  {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Status %s\n", resp.Status)

	if _, err := io.Copy(out, resp.Body); err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "fetch: reading %s, %v\n", url, err)
		os.Exit(1)
	}
	resp.Body.Close()
}


