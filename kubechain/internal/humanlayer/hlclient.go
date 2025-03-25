// Note, may eventually move on from the client.go in this project
// in which case I would rename this file to client.go
package humanlayer

import (
	"context"
	"os"

	humanlayerapi "github.com/humanlayer/smallchain/kubechain/internal/humanlayerapi"
)

// NewHumanLayerClient creates a new API client using either the provided API key
// or falling back to the HUMANLAYER_API_KEY environment variable
func NewHumanLayerClient(ctx context.Context, apiKey string) *humanlayerapi.APIClient {
	config := humanlayerapi.NewConfiguration()
	config.Host = "api.humanlayer.dev"
	config.Scheme = "https"

	// Use provided apiKey if not empty, otherwise fall back to environment variable
	key := apiKey
	if key == "" {
		key = os.Getenv("HUMANLAYER_API_KEY")
	}

	config.DefaultHeader["Authorization"] = "Bearer " + key

	client := humanlayerapi.NewAPIClient(config)

	return client
}
