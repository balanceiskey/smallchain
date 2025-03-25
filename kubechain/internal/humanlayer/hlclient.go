// Note, may eventually move on from the client.go in this project
// in which case I would rename this file to client.go
package humanlayer

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	humanlayerapi "github.com/humanlayer/smallchain/kubechain/internal/humanlayerapi"
)

// NewHumanLayerClient creates a new API client using either the provided API key
// or falling back to the HUMANLAYER_API_KEY environment variable. Similarly,
// it uses the provided API base URL or falls back to HUMANLAYER_API_BASE.
func NewHumanLayerClient(optionalApiBase string, optionalApiKey ...string) (*humanlayerapi.APIClient, error) {
	config := humanlayerapi.NewConfiguration()

	// Get API base from parameter or environment variable
	apiBase := os.Getenv("HUMANLAYER_API_BASE")
	if optionalApiBase != "" {
		apiBase = optionalApiBase
	}

	if apiBase == "" {
		apiBase = "https://api.humanlayer.dev/humanlayer/v1"
	}

	parsedURL, err := url.Parse(apiBase)
	if err != nil {
		return nil, fmt.Errorf("failed to parse API base URL: %v", err)
	}

	config.Host = strings.TrimPrefix(parsedURL.Host+parsedURL.Path, "/")
	config.Scheme = parsedURL.Scheme

	// Get API key from environment variable or optional parameter
	apiKey := os.Getenv("HUMANLAYER_API_KEY")
	if len(optionalApiKey) > 0 && optionalApiKey[0] != "" {
		apiKey = optionalApiKey[0]
	}

	if apiKey == "" {
		return nil, fmt.Errorf("API key not found in environment variable HUMANLAYER_API_KEY or parameter")
	}

	// Add authorization header to all requests
	config.DefaultHeader["Authorization"] = "Bearer " + apiKey

	// Create the API client with the configuration
	client := humanlayerapi.NewAPIClient(config)

	return client, nil
}
