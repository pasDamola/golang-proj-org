package main

import "net/http"

func main() {
	server := http.Server{
		Addr:    ":8181",
		Handler: http.FileServer(http.Dir("web/static")),
	}
	server.ListenAndServe()
}
