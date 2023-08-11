package main

import (
	"context"
	"time"

	"github.com/tetratelabs/log"
	"github.com/tetratelabs/telemetry"
	"github.com/tetratelabs/telemetry/scope"

	"github.com/mt-inside/tet-log-test/pkg/one"
	"github.com/mt-inside/tet-log-test/pkg/server"
	"github.com/mt-inside/tet-log-test/pkg/two"
)

func main() {
	log := log.NewFlattened()
	scope.UseLogger(log)
	scope.SetAllScopes(telemetry.LevelInfo)

	one.One()

	two := two.NewTwo(42, "hello")
	two.One()
	two.Two()

	s := server.NewServer()
	deadline, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s.HandleRequest(deadline, map[string]string{"path": "/foo", "data": "bar"})
}
