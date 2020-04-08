package main

import (
	"net/http"

	_ "net/http/pprof"
)

func main() {
	// only listening for pprof
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
}
