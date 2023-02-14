package sendpulse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	Auth   *AuthService
	Email  *EmailService
	Events *EventService

	baseURL       string
	eventsBaseURL string
	client        *http.Client
	config        *ClientConfig
}

type ClientConfig struct {
	ClientID     string
	ClientSecret string
}

func New(config *ClientConfig) *Client {
	c := &Client{
		baseURL:       "https://api.sendpulse.com",
		eventsBaseURL: "https://events.sendpulse.com",
		client: &http.Client{
			Timeout: time.Second * 30,
		},
		config: config,
	}

	c.Auth = newAuthService(c)
	c.Email = newEmailService(c)
	c.Events = newEventService(c)

	return c
}

func (c *Client) Send(method, path string, body, result any, useToken bool) (*http.Response, error) {
	return c.send(method, c.baseURL+path, body, result, useToken)
}

func (c *Client) InvokeEvent(method, path string, body, result any) (*http.Response, error) {
	return c.send(method, c.eventsBaseURL+path, body, result, true)
}

func (c *Client) send(method, url string, body, result any, useToken bool) (*http.Response, error) {
	if useToken && time.Now().After(c.Auth.Token.ExpiresAt) {
		if _, err := c.Auth.Authorize(c.config.ClientID, c.config.ClientSecret); err != nil {
			return nil, err
		}
	}

	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		r = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, url, r)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "sendpulse-go-http-client")
	req.Header.Set("Accept", "*/*")

	if useToken {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Auth.Token.Value))
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		//var apiErr APIError
		apiErr := NewAPIError(res.StatusCode, req.URL.String(), string(resBody), nil)
		if err := json.Unmarshal(resBody, apiErr); err != nil {
			return nil, apiErr
		}

		return nil, apiErr
	}

	if result != nil {
		if err := json.Unmarshal(resBody, result); err != nil {
			return nil, NewAPIError(res.StatusCode, req.URL.String(), string(resBody), err)
		}
	}

	return res, nil
}
