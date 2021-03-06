package main

import (
	"sync"
	"net/http"
	"log"
	"fmt"
)

var mu sync.Mutex
var count int

func main()  {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler1(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w, "Counter = %d\n", count)
	mu.Unlock()
}
