// An addition of the counter functionalities to the webserver.go
//NOTE
// Check wserver.go file for instruction on how to run this locally
// Assess counter functionality via "localhost:8080/count"

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path: %q\n", r.URL.Path)
}

// enter "localhost:8080/count" in any web browser on your system
//get the number of times requests has been made
func counter(write http.ResponseWriter, pth *http.Request) {
	mu.Lock()
	fmt.Fprintf(write, "count: %d\n", count)
	mu.Unlock()
}
