package main

import (
	"os"
	"net/http"
	"fmt"
)

const (
	defaultAddr = ":8888"
)

func main() {
	var addr string
	dir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	if len(os.Args) > 1 {
		addr = os.Args[1]
	} else {
		addr = defaultAddr
	}

	fmt.Printf("Serving files from %s over http on %s\n", dir, addr)
	http.Handle("/", http.FileServer(http.Dir(dir)))
	http.ListenAndServe(addr, nil)
}

