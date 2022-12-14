//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

// wire.go

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
