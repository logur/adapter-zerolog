# Logur adapter for [zerolog](https://github.com/rs/zerolog)

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/logur/adapter-zerolog/CI?style=flat-square)
[![Codecov](https://img.shields.io/codecov/c/github/logur/adapter-zerolog?style=flat-square)](https://codecov.io/gh/logur/adapter-zerolog)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/adapter/zerolog?style=flat-square)](https://goreportcard.com/report/logur.dev/adapter/zerolog)
[![GolangCI](https://golangci.com/badges/github.com/logur/adapter-zerolog.svg)](https://golangci.com/r/github.com/logur/adapter-zerolog)
[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-61CFDD.svg?style=flat-square)](https://github.com/logur/adapter-zerolog)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/logur.dev/adapter/zerolog)


## Installation

```bash
go get logur.dev/adapter/zerolog
```


## Usage

```go
package main

import (
	"github.com/rs/zerolog"
	zerologadapter "logur.dev/adapter/zerolog"
)

func main() {
	logger := zerologadapter.New(zerolog.Nop())
}
```


## Development

When all coding and testing is done, please run the test suite:

```bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
