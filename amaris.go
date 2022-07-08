package amarisapi

import (
	"log"
	amaris "testamaris/gen/amaris"
)

// amaris service example implementation.
// The example methods log the requests and return zero values.
type amarissrvc struct {
	logger *log.Logger
}

// NewAmaris returns the amaris service implementation.
func NewAmaris(logger *log.Logger) amaris.Service {
	return &amarissrvc{logger}
}
