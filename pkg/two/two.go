package two

import (
	"github.com/tetratelabs/telemetry"
	"github.com/tetratelabs/telemetry/scope"
)

var log = scope.Register("two", "package two")

type Two struct {
	log  telemetry.Logger
	ID   int32
	Data string
}

func NewTwo(id int32, data string) *Two {
	return &Two{
		log.With("id", id),
		id,
		data,
	}
}

func (t *Two) One() {
	t.log.Info("step 1", "data len", len(t.Data))
}
func (t *Two) Two() {
	t.log.Info("step 2")
}
