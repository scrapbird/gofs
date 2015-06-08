package main

import (
	"os"
	"net/http"
	"fmt"
	"strconv"
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
		// check address format
		c := string([]rune(addr)[0])
		if _, err := strconv.Atoi(c); err == nil {
			if c != ":" {
				addr = ":" + addr
			}
		}
	} else {
		addr = defaultAddr
	}

	fmt.Printf("Serving files from %s over http on %s\n", dir, addr)
	http.Handle("/", http.FileServer(http.Dir(dir)))
	http.ListenAndServe(addr, nil)
}

