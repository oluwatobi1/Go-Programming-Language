// A simple web server with go

// To how run this locally
// Make sure you have ...
// Golang install locally and added to system path (i.e Golang commands can be access from terminal)
// Get the directory of the location of this file (folder path)
// Open the command terminal and navigate there(i.e on window-- cd C:\Users\HP\Desktop\Go\webServer2)
// Enter 'go build wserver.go' ie wserver.go is this filename (the command will not return an output)
// Open a local web browser and enter 'localhost:8080'

package main

import (
	"fmt"
	"log"
	"net/http"
)

//uses a handler function
//Note it is easier in go cos of the available library functions
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path: %q\n", r.URL.Path)
}
