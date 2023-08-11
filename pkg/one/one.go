package one

import "github.com/tetratelabs/telemetry/scope"

var log = scope.Register("one", "the one package")

func One() {
	log.Info("step 1")
	OneSub()
}
