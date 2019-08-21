package zerolog_test

import (
	zerologadapter "logur.dev/adapter/zerolog"
)

func ExampleNew() {
	var l interface{}
	logger := zerologadapter.New(l)

	// Output:
	_ = logger
}

// If logger is nil, a default instance is created.
func ExampleNew_default() {
	logger := zerologadapter.New(nil)

	// Output:
	_ = logger
}
