// Code generated by goa v3.14.0, DO NOT EDIT.
//
// app gRPC client
//
// Command:
// $ goa gen app/design

package client

import (
	apppb "app/gen/grpc/app/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli apppb.AppClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the app service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: apppb.NewAppClient(cc),
		opts:    opts,
	}
}

// Select calls the "Select" function in apppb.AppClient interface.
func (c *Client) Select() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildSelectFunc(c.grpccli, c.opts...),
			EncodeSelectRequest,
			DecodeSelectResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}