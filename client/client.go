// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/getzep/zep-go/v2/core"
	document "github.com/getzep/zep-go/v2/document"
	memory "github.com/getzep/zep-go/v2/memory"
	option "github.com/getzep/zep-go/v2/option"
	user "github.com/getzep/zep-go/v2/user"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Document *document.Client
	Memory   *memory.Client
	User     *user.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	if options.APIKey == "" {
		options.APIKey = os.Getenv("ZEP_API_KEY")
	}
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:   options.ToHeader(),
		Document: document.NewClient(opts...),
		Memory:   memory.NewClient(opts...),
		User:     user.NewClient(opts...),
	}
}
