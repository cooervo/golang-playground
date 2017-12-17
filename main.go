package main

import (
	"net/http"
	"time"
	"quoteapp/router"
)

func main() {
	println("starting server")

	mux := http.NewServeMux()

	// All requests that start with /static/* get handled by Handler created by StripPrefix
	fs := http.FileServer(http.Dir("public"))
	handler := http.StripPrefix("/static/", fs)
	mux.Handle("/static/", handler)

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
