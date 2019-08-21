package zerolog_test

import (
	"github.com/rs/zerolog"

	zerologadapter "logur.dev/adapter/zerolog"
)

func ExampleNew() {
	logger := zerologadapter.New(zerolog.Nop())

	// Output:
	_ = logger
}
