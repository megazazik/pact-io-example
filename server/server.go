package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func StartServer(outerCtx context.Context) {
	handler := http.NewServeMux()

	handler.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Sleep starting")
		time.Sleep(time.Second * 1)
		fmt.Println("Sleep end")
		fmt.Fprintf(w, "hello\n")
	})

	handler.HandleFunc("/counter", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"status\": \"ok\"}")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	stopServerChan := make(chan interface{})

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("Server stopped with error", err)
		}
		fmt.Println("After server stopped")
		stopServerChan <- nil
	}()

	fmt.Println("Server started")

	select {
	case <-stopServerChan:
		fmt.Println("Catch stopper server")
	case <-outerCtx.Done():
		fmt.Println("Catch outer context done")
	}

	shutDownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := server.Shutdown(shutDownCtx); err != nil {
		fmt.Println("Shutdown with error", err)
	}

	fmt.Println("Exit server")
}
