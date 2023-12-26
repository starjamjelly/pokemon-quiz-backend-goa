package appapi

import (
	app "app/gen/app"
	"context"
	"log"
)

// app service example implementation.
// The example methods log the requests and return zero values.
type appsrvc struct {
	logger *log.Logger
}

// NewApp returns the app service implementation.
func NewApp(logger *log.Logger) app.Service {
	return &appsrvc{logger}
}

// Select partner pokeomon by type.
func (s *appsrvc) Select(ctx context.Context, p *app.SelectPayload) (res string, err error) {
	switch p.Type {
	case "fire":
		res = "Charmander"
	case "water":
		res = "Squirtle"
	case "grass":
		res = "Bulbasaur"
	default:
		res = p.Type
	}
	return res, nil
}
