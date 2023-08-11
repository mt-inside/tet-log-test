package server

import (
	"context"
	"errors"

	"github.com/tetratelabs/telemetry"
	"github.com/tetratelabs/telemetry/scope"
)

var log = scope.Register("server", "my toy server")

type Server struct {
	resourceID uint32
	log        telemetry.Logger
}

func NewServer() *Server {
	resourceID := 69 // pretend we hold a db connection or something

	return &Server{
		resourceID: 69,
		log:        log.With("resource_id", resourceID),
	}
}

func (s *Server) HandleRequest(ctx context.Context, payload map[string]string) {
	ctx = telemetry.KeyValuesToContext(ctx, "requestPath", payload["path"])

	log := s.log.Context(ctx)

	log.Info("Handling request")

	s.processMore(ctx, payload)
}

func (s *Server) processMore(ctx context.Context, payload map[string]string) {
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return
	}

	log := s.log.Context(ctx)

	log.Info("Still processing")
}
