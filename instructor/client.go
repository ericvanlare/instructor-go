package instructor

import (
	"context"
	"encoding/json"
)

type JSONMarshalable[T any] interface {
	json.Marshaler
	json.Unmarshaler
}

type Client[T JSONMarshalable[T]] interface {
	CreateCompletion(
		ctx context.Context,
		messages []Message,
		opts ...ClientOptions,
	) (T, error)

	CreateChatCompletion(
		ctx context.Context,
		messages []Message,
		opts ...ClientOptions,
	) (T, error)

	// TODO: implement streaming
	// CreateCompletionStream() ()
	// CreateChatCompletionStream() ()
}

type ClientOptions struct {
	Mode       Mode
	MaxRetries int

	// Provider specific options:
}

func WithMode(mode Mode) ClientOptions {
	return ClientOptions{Mode: mode}
}

func WithMaxRetries(maxRetries int) ClientOptions {
	return ClientOptions{MaxRetries: maxRetries}
}

func FromOpenAI[T JSONMarshalable[T]]() *Client[T] {
	panic("not implemented")
}