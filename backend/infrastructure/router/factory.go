package router

import (
	"errors"
	"time"

	"github.com/h4shu/lounge-books/application/repositories"
)

type Server interface {
	Listen()
}

type Port int

var (
	errInvalidWebServerInstance = errors.New("invalid router server instance")
)

const (
	InstanceEcho int = iota
)

func NewWebServerFactory(
	instance int,
	dbSQL repositories.SQL,
	port Port,
	ctxTimeout time.Duration,
) (Server, error) {
	switch instance {
	case InstanceEcho:
		return newEchoServer(dbSQL, port, ctxTimeout), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
