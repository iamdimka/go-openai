package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/iamdimka/go-openai"
)

func example(ctx context.Context) error {
	var (
		org = flag.String("org", "", "OpenAI organization")
		key = flag.String("key", "", "OpenAI API key")
	)
	flag.Parse()

	c := openai.NewClient(*org, *key)

	model, err := c.Model(ctx, "gpt-3.5-turbo")
	if err != nil {
		return err
	}

	fmt.Printf("Model: %+v\n", model)

	err = c.ChatCompletionsStream(ctx, openai.ChatCompletionsRequest{
		Model:       "gpt-3.5-turbo",
		Temperature: 0.7,
		Messages: []openai.Message{
			{"user", "Hello, how are you?"},
		},
	}, func(c *openai.ChatCompletionChunk) error {
		_, err := fmt.Printf(c.Choices[0].Delta.Content)
		return err
	})
	return err
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
		case <-s:
		}

		signal.Stop(s)
		close(s)
		cancel()
	}()

	err := example(ctx)
	cancel()
	if err != nil {
		panic(err)
	}
}
