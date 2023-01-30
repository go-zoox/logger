package console

import (
	"fmt"
	"log"

	"github.com/go-zoox/logger/components/constants"
	"github.com/go-zoox/logger/components/transport"
)

type Console struct {
	level  string
	logger *log.Logger
}

type Config struct {
	Level string
}

func New(config ...*Config) transport.Transport {
	level := constants.LevelDebug
	if len(config) > 0 {
		if config[0].Level != "" {
			level = config[0].Level
		}
	}

	return &Console{
		level: level,
	}
}

func (c *Console) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}
