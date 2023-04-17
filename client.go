package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client is the OpenAI API client.
type Client struct {
	Organization string
	APIKey       string

	authorization string
}

// NewClient creates a new OpenAI API client.
func NewClient(organization, apiKey string) *Client {
	return &Client{
		Organization: organization,
		APIKey:       apiKey,

		authorization: "Bearer " + apiKey,
	}
}

// Request makes a request to the OpenAI API.
func (c *Client) Request(ctx context.Context, path string, payload, result any) error {
	var (
		method string
		body   io.ReadWriter
		err    error
	)

	if payload == nil {
		method = http.MethodGet
	} else {
		method = http.MethodPost
		body = bytes.NewBuffer(nil)
		err = json.NewEncoder(body).Encode(payload)
		if err != nil {
			return err
		}
	}

	if ctx == nil {
		ctx = context.Background()
	}

	req, err := http.NewRequestWithContext(ctx, method, "https://api.openai.com/v1/"+path, body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", c.authorization)
	req.Header.Set("OpenAI-Organization", c.Organization)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var msg struct {
			Error ApiError `json:"error"`
		}

		err = json.NewDecoder(res.Body).Decode(&msg)
		if err != nil {
			return fmt.Errorf("got status code %d", res.StatusCode)
		}

		return msg.Error
	}

	if dec, ok := result.(Handler); ok {
		return dec.Handle(ctx, res)
	}

	return json.NewDecoder(res.Body).Decode(result)
}
