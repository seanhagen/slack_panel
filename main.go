package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/**
 * File: main.go
 * Date: 2021-08-13 14:38:43
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/").Handler(fs)
	http.Handle("/", r)
	fmt.Printf("Starting server on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Panic(err)
	}
}
