package sendpulse

import "time"

type AuthService struct {
	Token *AuthToken

	client *Client
}

func newAuthService(client *Client) *AuthService {
	return &AuthService{
		Token:  &AuthToken{},
		client: client,
	}
}

type AuthParams struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type Auth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (s *AuthService) Authorize(clientID, clientSecret string) (*Auth, error) {
	params := AuthParams{
		GrantType:    "client_credentials",
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	auth := &Auth{}
	if _, err := s.client.Send("POST", "/oauth/access_token", params, auth, false); err != nil {
		return nil, err
	}

	s.Token = &AuthToken{
		Value:     auth.AccessToken,
		ExpiresAt: time.Now().Add(time.Second * time.Duration(auth.ExpiresIn-10)),
	}

	return auth, nil
}

type AuthToken struct {
	Value     string
	ExpiresAt time.Time
}
