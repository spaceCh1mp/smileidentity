// client to consume smileID API endpoints
package smileidentity

type Client struct {
	config *ClientConfig
}

type ClientConfig struct {
	smilePartnerID string
	smileAPIKey    string

	// prod directs requests to the smile production environment if set to true.
	// prod is false by default
	prod bool
}

func NewClient(config *ClientConfig) *Client {
	return &Client{
		config: config,
	}
}
