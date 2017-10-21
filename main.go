package main

import (
	"net/http"
	"time"
	"golang-playground/router"
)

func main() {
	println("starting server")

	mux := http.NewServeMux()

	// Serve public files
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", router.Index)
	mux.HandleFunc("/documentary", router.Documentary)

	server := &http.Server{
		Addr:         config.Address,
		Handler:      mux,
		ReadTimeout:  time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout: time.Duration(config.WriteTimeout * int64(time.Second)),
	}

	server.ListenAndServe()
	println("Server running at http://" + config.Address)

}
