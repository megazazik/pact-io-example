package server_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/megazazik/pact-io-example/server"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestStartServer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go server.StartServer(ctx)

	defer func() {
		cancel()
	}()

	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Provider: "MyProvider",
	}

	// Verify the Provider using the locally saved Pact Files
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:8080",
		PactURLs:        []string{filepath.ToSlash("../pacts/metricsclient-metricsapimiddleware.json")},
		// StateHandlers: types.StateHandlers{
		//   // Setup any state required by the test
		//   // in this case, we ensure there is a "user" in the system
		//   "User foo exists": func() error {
		// 	lastName = "crickets"
		// 	return nil
		//   },
		// },
	})

	if err != nil {
		t.Fatalf("Error during tests: %v", err)
	}
}
