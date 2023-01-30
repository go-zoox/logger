package file

import (
	"log"
	"os"

	"github.com/go-zoox/logger/components/constants"
	"github.com/go-zoox/logger/components/transport"
)

type File struct {
	level string
	// file   string
	logger *log.Logger
}

type Config struct {
	Level string
	File  string
}

func New(config ...*Config) transport.Transport {
	level := constants.LevelDebug
	var logger *log.Logger
	if len(config) > 0 {
		if config[0].Level != "" {
			level = config[0].Level
		}
		if config[0].File != "" {
			file, err := os.OpenFile(config[0].File, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				panic("error opening file: " + config[0].File + ": " + err.Error())
			}
			// defer file.Close()
			logger = log.New(file, "", log.Ldate|log.Ltime)
		}
	}

	return &File{
		level:  level,
		logger: logger,
	}
}

func (f *File) Write(p []byte) (n int, err error) {
	if f.logger == nil {
		return
	}

	f.logger.Println(string(p))

	return len(p), nil
}
