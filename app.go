package main

import (
	"context"
	"time"

	"github.com/megazazik/pact-io-example/server"
)

func main() {
	shutDownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	server.StartServer(shutDownCtx)
}
