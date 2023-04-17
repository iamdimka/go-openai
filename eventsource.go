package openai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	dataPrefix = []byte("data: ")
)

// EventSource is a streaming event source.
type EventSource[T any] struct {
	handler func(*T) error
}

// NewEventSource creates a new EventSource.
func NewEventSource[T any](handler func(d *T) error) *EventSource[T] {
	return &EventSource[T]{
		handler: handler,
	}
}

// Decode decodes the response into events.
func (e *EventSource[T]) Handle(ctx context.Context, res *http.Response) error {
	var (
		data   []byte
		reader = bufio.NewReader(res.Body)
	)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			return err
		}

		if len(line) == 1 {
			if data != nil {
				if isDoneMessage(data) {
					return nil
				}

				err = e.unmarshalAndFlush(data)
				if err != nil {
					return err
				}

				data = nil
			}

			continue
		}

		if !bytes.HasPrefix(line, dataPrefix) {
			return fmt.Errorf("invalid line: %q", line)
		}

		data = append(data, line[len(dataPrefix):len(line)-1]...)
	}
}

// isDoneMessage checks if the data is a done message.
func isDoneMessage(data []byte) bool {
	return len(data) == 6 && string(data) == "[DONE]"
}

// unmarshalAndFlush unmarshals the data and flushes it to the handler.
func (e *EventSource[T]) unmarshalAndFlush(data []byte) error {
	v := new(T)
	err := json.Unmarshal(data, v)
	if err == nil {
		return e.handler(v)
	}

	return err
}
