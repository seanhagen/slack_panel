package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/seanhagen/slack_panel/backend"
)

/**
 * File: main.go
 * Date: 2021-08-13 14:38:43
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

func main() {
	hub := backend.NewHub()
	go hub.Run()

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		backend.ServeWs(hub, w, r)
	})

	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/").Handler(fs)

	fmt.Println("Starting server on port 3000")
	l := handlers.LoggingHandler(os.Stdout, r)
	if err := http.ListenAndServe(":3000", l); err != nil {
		log.Panic(err)
	}
}
