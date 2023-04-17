package openai

import "context"

// EditRequest is the request for the Edit API.
type EditRequest struct {
	Model       string  `json:"model"`
	Input       string  `json:"input,omitempty"`
	Instruction string  `json:"instruction"`
	N           int     `json:"n,omitempty"`
	Temperature float32 `json:"temperature,omitempty"`
	TopP        float32 `json:"top_p,omitempty"`
}

// Edit is the response from the Edit API.
type Edit struct {
	Object  string       `json:"object"`
	Created int64        `json:"created"`
	Choices []EditChoice `json:"choices"`
	Usage   Usage        `json:"usage"`
}

// EditChoice is a single choice from the Edit API.
type EditChoice struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

// Edit calls the Edit API.
func (c *Client) Edit(ctx context.Context, r EditRequest) (*Edit, error) {
	res := new(Edit)

	err := c.Request(ctx, "edits", nil, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
