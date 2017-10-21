package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static", http.StripPrefix("/static/", files))

	//TODO mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:         config.Address,
		Handler:      mux,
		ReadTimeout:  time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout: time.Duration(config.WriteTimeout * int64(time.Second)),
	}

	server.ListenAndServe()
	print("Server running at ", "http://"+config.Address)

}
