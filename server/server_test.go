package main

import (
	"fmt"
	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../pacts", dir)

// Start provider API in the background
func startServer() {
	go startProvider()
	time.Sleep(2 * time.Second)
}

func TestServerPact_Verification(t *testing.T) {

	startServer()

	verifier := provider.NewVerifier()
	err := verifier.VerifyProvider(t, provider.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:8081",
		Provider:        "BookProvider",
		PactFiles: []string{
			filepath.ToSlash(fmt.Sprintf("%s/SampleBookConsumer-SampleBookProvider.json", pactDir)),
		},
		DisableColoredOutput: true,
	})

	assert.NoError(t, err)
}
