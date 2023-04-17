package openai

import (
	"context"
	"net/http"
)

type ApiError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (e ApiError) Error() string {
	return e.Message
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Handler interface {
	Handle(context.Context, *http.Response) error
}
