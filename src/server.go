package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := &http.ServeMux{}

	handler.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello\n")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	fmt.Println("Server started")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server stopped with error", err)
	}
}
