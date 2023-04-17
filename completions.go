package openai

import "context"

// CompletionsRequest is a request to the Completions API.
type CompletionsRequest struct {
	Model            string         `json:"model"`
	Prompt           string         `json:"prompt"`
	Suffix           string         `json:"suffix,omitempty"`
	MaxTokens        int            `json:"max_tokens"`
	Temperature      float32        `json:"temperature,omitempty"`
	TopP             float32        `json:"top_p,omitempty"`
	N                int            `json:"n,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	Logprobs         int            `json:"logprobs,omitempty"`
	Echo             bool           `json:"echo,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	PresencePenalty  float32        `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32        `json:"frequency_penalty,omitempty"`
	BestOf           int            `json:"best_of,omitempty"`
	LogitBias        map[string]any `json:"logit_bias,omitempty"`
	User             string         `json:"user,omitempty"`
}

// Validate validates the CompletionsRequest.
func (c *CompletionsRequest) Validate() error {
	return nil
}

// TextCompletion is the response from the Completions API.
type TextCompletion struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Choices []TextCompletionChoice `json:"choices"`
	Usage   Usage                  `json:"usage"`
}

// TextCompletionChunk is the response from the Completions API.
type TextCompletionChunk struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Choices []TextCompletionChoice `json:"choices"`
}

// TextCompletionChoice is a single choice from the Completions API.
type TextCompletionChoice struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	Logprobs     any    `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

// Completions returns a text completion.
func (c *Client) Completions(ctx context.Context, req CompletionsRequest) (*TextCompletion, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	req.Stream = false
	res := new(TextCompletion)
	err = c.Request(ctx, "completions", req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// CompletionsStream returns a text completion.
func (c *Client) CompletionsStream(ctx context.Context, req CompletionsRequest, cb func(*TextCompletionChunk) error) error {
	err := req.Validate()
	if err != nil {
		return err
	}

	req.Stream = true
	return c.Request(ctx, "completions", req, NewEventSource(cb))
}
