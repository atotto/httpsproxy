package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	director := func(request *http.Request) {
		request.URL.Scheme = "http"
		request.URL.Host = ":8080"
	}
	rp := &httputil.ReverseProxy{Director: director}
	server := http.Server{
		Addr:    ":8443",
		Handler: rp,
	}

	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
