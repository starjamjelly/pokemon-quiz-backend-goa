// Code generated by goa v3.14.0, DO NOT EDIT.
//
// app gRPC server types
//
// Command:
// $ goa gen app/design

package server

import (
	app "app/gen/app"
	apppb "app/gen/grpc/app/pb"
)

// NewSelectPayload builds the payload of the "select" endpoint of the "app"
// service from the gRPC request type.
func NewSelectPayload(message *apppb.SelectRequest) *app.SelectPayload {
	v := &app.SelectPayload{
		Type: message.Type,
	}
	return v
}

// NewProtoSelectResponse builds the gRPC response type from the result of the
// "select" endpoint of the "app" service.
func NewProtoSelectResponse(result string) *apppb.SelectResponse {
	message := &apppb.SelectResponse{}
	message.Field = result
	return message
}
