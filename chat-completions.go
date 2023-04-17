package openai

import "context"

type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

// ChatCompletionsRequest is the request for the chat completions endpoint.
type ChatCompletionsRequest struct {
	Model            string         `json:"model"`
	Messages         []Message      `json:"messages"`
	Temperature      float32        `json:"temperature,omitempty"`
	TopP             float32        `json:"top_p,omitempty"`
	N                int            `json:"n,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	PresencePenalty  float32        `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32        `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]any `json:"logit_bias,omitempty"`
	User             string         `json:"user,omitempty"`
}

// ChatCompletion is the response from the chat completions endpoint.
type ChatCompletion struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
	Choices []Choice `json:"choices"`
}

// ChatCompletionChunk is the response from the chat completions endpoint.
type ChatCompletionChunk struct {
	ID      string  `json:"id"`
	Object  string  `json:"object"`
	Created int64   `json:"created"`
	Model   string  `json:"model"`
	Choices []Delta `json:"choices"`
}

type Delta struct {
	Delta Message `json:"delta"`
}

// Choice is a single choice from the chat completions endpoint.
type Choice struct {
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
}

// Message is a single message from the chat completions endpoint.
type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

// ChatCompletions returns a chat completion.
func (c *Client) ChatCompletions(ctx context.Context, r ChatCompletionsRequest) (*ChatCompletion, error) {
	res := new(ChatCompletion)
	r.Stream = false

	if err := c.Request(ctx, "chat/completions", r, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ChatCompletionsStream returns a chat completion.
func (c *Client) ChatCompletionsStream(ctx context.Context, r ChatCompletionsRequest, cb func(*ChatCompletionChunk) error) error {
	r.Stream = true
	return c.Request(ctx, "chat/completions", r, NewEventSource(cb))
}
