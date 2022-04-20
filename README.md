# Logger

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/logger)](https://pkg.go.dev/github.com/go-zoox/logger)
[![Build Status](https://github.com/go-zoox/logger/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/logger/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/logger)](https://goreportcard.com/report/github.com/go-zoox/logger)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/logger/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/logger?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/logger.svg)](https://github.com/go-zoox/logger/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/logger.svg?label=Release)](https://github.com/go-zoox/logger/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/logger
```

## Getting Started

```go
logger.Debug("Hi, %s", "Goo Zoox")
logger.Info("Hi, %s", "Goo Zoox")
logger.Warn("Hi, %s", "Goo Zoox")
logger.Error("Hi, %s", "Goo Zoox")
logger.Fatal("Hi, %s", "Goo Zoox")

// set level
logger.SetLevel("error")
```

## Inspired by
* [kenshinx/godns](https://github.com/kenshinx/godns/blob/master/log.go) - About
A fast dns cache server written by go.
* [winstonjs/winston](https://github.com/winstonjs/winston) - A logger for just about everything.

## License
GoZoox is released under the [MIT License](./LICENSE).