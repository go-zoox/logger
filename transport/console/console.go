package console

import (
	"io"
	"log"
	"os"

	"github.com/go-zoox/logger/components/constants"
	"github.com/go-zoox/logger/components/transport"
)

type Console struct {
	level  string
	logger *log.Logger
	//
	stdout io.Writer
}

type Option struct {
	Level string
	//
	Stdout io.Writer
}

func New(opts ...func(opt *Option)) transport.Transport {
	opt := &Option{
		Level:  constants.LevelDebug,
		Stdout: os.Stdout,
	}
	for _, o := range opts {
		o(opt)
	}

	return &Console{
		level:  opt.Level,
		stdout: opt.Stdout,
	}
}

func (c *Console) Write(p []byte) (n int, err error) {
	line := append([]byte{}, p...)
	line = append(line, '\n')
	c.stdout.Write(line)
	return len(p), nil
}

func (c *Console) WriteWithLevel(p []byte, level string) (n int, err error) {
	return c.Write(p)
}
