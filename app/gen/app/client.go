// Code generated by goa v3.14.0, DO NOT EDIT.
//
// app client
//
// Command:
// $ goa gen app/design

package app

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "app" service client.
type Client struct {
	SelectEndpoint goa.Endpoint
}

// NewClient initializes a "app" service client given the endpoints.
func NewClient(select_ goa.Endpoint) *Client {
	return &Client{
		SelectEndpoint: select_,
	}
}

// Select calls the "select" endpoint of the "app" service.
func (c *Client) Select(ctx context.Context, p *SelectPayload) (res string, err error) {
	var ires any
	ires, err = c.SelectEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}
