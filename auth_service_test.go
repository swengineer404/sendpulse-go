package sendpulse

import "testing"

func TestAuthService_Authorize(t *testing.T) {
	client := newTestClient()
	auth, err := client.Auth.Authorize(client.config.ClientID, client.config.ClientSecret)
	if err != nil {
		t.Fatal(err)
	}

	if auth.AccessToken == "" {
		t.Fatal("auth token not found")
	}
}
