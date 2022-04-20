package console

import (
	"log"
	"os"

	"github.com/go-zoox/logger/components/constants"
	"github.com/go-zoox/logger/components/message"
	"github.com/go-zoox/logger/components/transport"
)

type Console struct {
	level  int
	logger *log.Logger
}

type Config struct {
	Level int
}

func New(config ...*Config) transport.Transport {
	level := constants.LevelDebug
	if len(config) > 0 {
		if config[0].Level > 0 {
			level = config[0].Level
		}
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	return &Console{
		level:  level,
		logger: logger,
	}
}

func (c *Console) Write(message *message.Message) {
	if c.level <= message.Level {
		c.logger.Println(message.Message)
	}
}
