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
	//
	cfg *Config
}

type Config struct {
	FilePath string
	//
	File *os.File
	// Level specify the min level.
	Level string
	// Extact means only the specify level to write
	// Which is used for custom access, error, debug log.
	Exact bool
}

func New(cfg ...*Config) transport.Transport {
	var err error
	cfgX := &Config{
		Level: constants.LevelDebug,
	}

	if len(cfg) > 0 && cfg[0] != nil {
		cfgX = cfg[0]
	}

	if cfg[0].FilePath != "" {
		cfgX.File, err = os.OpenFile(cfg[0].FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic("error opening file: " + cfg[0].FilePath + ": " + err.Error())
		}
	}

	if cfgX.File == nil {
		panic("file transport need a file, but not provided")
	}

	return &File{
		cfg:    cfgX,
		level:  cfgX.Level,
		logger: log.New(cfgX.File, "", log.Ldate|log.Ltime),
	}
}

func (f *File) Write(p []byte) (n int, err error) {
	f.logger.Println(string(p))

	return len(p), nil
}

func (f *File) WriteWithLevel(p []byte, level string) (n int, err error) {
	if !f.cfg.Exact {
		f.logger.Println(string(p))
	} else {
		if f.level == level {
			f.logger.Println(string(p))
		}
	}

	return len(p), nil
}
