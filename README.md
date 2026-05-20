# Logger

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/logger)](https://pkg.go.dev/github.com/go-zoox/logger)
[![Build Status](https://github.com/go-zoox/logger/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/logger/actions/workflows/ci.yml)

## Installation

```bash
go get github.com/go-zoox/logger
```

## Usage

```go
logger.Info("hello")
logger.Error("something failed")
```

## Transports

A `Logger` holds named **transports**. Each log line is sent to **every** transport; each transport decides what to do (console, files, etc.).

### Console (default)

```go
l := logger.New() // includes console transport
```

### File transport with per-level paths

`file.Config.FilePath` is the **default** file. Optional `Levels` sends specific log levels to other files; all other levels still use `FilePath`.

```go
import (
	"github.com/go-zoox/logger"
	"github.com/go-zoox/logger/transport/console"
	"github.com/go-zoox/logger/transport/file"
)

l := logger.New(func(opt *logger.Option) {
	opt.Transports = map[string]transport.Transport{
		"console": console.New(),
		"file": file.New(&file.Config{
			FilePath: "/var/log/app.log",
			Levels: map[string]string{
				"error": "/var/log/error.log",
			},
		}),
	}
})

l.Info("→ console + app.log")
l.Error("→ console + error.log")
```

Omit `Levels` (or leave it empty) to write every line to `FilePath` only. Level keys: `debug`, `info`, `warn`, `error`, `fatal`.

### Multiple transports

Use several entries in `opt.Transports` (e.g. `console` + `file` above). No extra sink/output layer.

## License

[MIT](./LICENSE)
