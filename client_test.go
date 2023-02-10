package sendpulse

import "os"

func newTestClient() *Client {
	config := &ClientConfig{
		ClientID:     os.Getenv("SP_CLIENT_ID"),
		ClientSecret: os.Getenv("SP_CLIENT_SECRET"),
	}
	return New(config)
}
