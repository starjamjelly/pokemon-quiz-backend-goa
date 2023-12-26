// Code generated by goa v3.14.0, DO NOT EDIT.
//
// app service
//
// Command:
// $ goa gen app/design

package app

import (
	"context"
)

// The app service provides some pokemon api.
type Service interface {
	// Select implements select.
	Select(context.Context, *SelectPayload) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "app"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"select"}

// SelectPayload is the payload type of the app service select method.
type SelectPayload struct {
	// Selected Pokemon Types
	Type string
}
