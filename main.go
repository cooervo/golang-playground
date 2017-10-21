package main

import (
	"net/http"
	"time"
	"golang-playground/router"
)

func main() {
	print("starting server")

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", router.Index)

	server := &http.Server{
		Addr:         config.Address,
		Handler:      mux,
		ReadTimeout:  time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout: time.Duration(config.WriteTimeout * int64(time.Second)),
	}

	server.ListenAndServe()
	print("Server running at http://" + config.Address)

}
