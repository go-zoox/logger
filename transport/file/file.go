package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/go-zoox/logger/components/constants"
	"github.com/go-zoox/logger/components/transport"
)

type File struct {
	cfg *Config
	// single-file mode (Levels empty)
	level  string
	logger *log.Logger
	// level-routing mode (Levels non-empty)
	files   map[string]string
	mu      sync.Mutex
	loggers map[string]*log.Logger
}

// Config configures the file transport.
type Config struct {
	// FilePath is the default log file. Used for every line when Levels is empty,
	// or for levels not listed in Levels when Levels is set.
	FilePath string
	// File is an optional open file handle (single-file mode only).
	File *os.File
	// Level is the min level when Exact is true (legacy).
	Level string
	// Exact means only the specified Level is written (legacy single-file mode).
	Exact bool
	// Levels maps log level names (debug, info, warn, error, fatal) to file paths.
	// Unlisted levels use FilePath. When empty, all lines go to FilePath only.
	Levels map[string]string
}

func New(cfg ...*Config) transport.Transport {
	cfgX := &Config{
		Level: constants.LevelDebug,
	}
	if len(cfg) > 0 && cfg[0] != nil {
		cfgX = cfg[0]
		cfgX.Level = strings.ToUpper(strings.TrimSpace(cfgX.Level))
	}

	if len(cfgX.Levels) > 0 {
		return newLevelFiles(cfgX)
	}
	return newSingleFile(cfgX)
}

func newSingleFile(cfg *Config) transport.Transport {
	var err error
	if cfg.File == nil && cfg.FilePath != "" {
		cfg.File, err = os.OpenFile(cfg.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
		if err != nil {
			panic("error opening file: " + cfg.FilePath + ": " + err.Error())
		}
	}
	if cfg.File == nil {
		panic("file transport need a file, but not provided")
	}
	return &File{
		cfg:    cfg,
		level:  cfg.Level,
		logger: log.New(cfg.File, "", log.Ldate|log.Ltime),
	}
}

func newLevelFiles(cfg *Config) transport.Transport {
	if strings.TrimSpace(cfg.FilePath) == "" && len(cfg.Levels) == 0 {
		panic("file transport: FilePath or Levels is required")
	}
	files := make(map[string]string, len(cfg.Levels)+1)
	for k, path := range cfg.Levels {
		key := normalizeLevelKey(k)
		path = strings.TrimSpace(path)
		if path == "" {
			panic(fmt.Sprintf("file transport: empty path for level %q", k))
		}
		mkdirForLogFile(path)
		files[key] = path
	}
	if p := strings.TrimSpace(cfg.FilePath); p != "" {
		mkdirForLogFile(p)
		files["default"] = p
	}
	return &File{
		cfg:     cfg,
		loggers: make(map[string]*log.Logger),
		files:   files,
	}
}

func normalizeLevelKey(k string) string {
	k = strings.ToLower(strings.TrimSpace(k))
	switch k {
	case "warning":
		return "warn"
	case "err":
		return "error"
	default:
		return k
	}
}

func mkdirForLogFile(path string) {
	dir := filepath.Dir(path)
	if dir == "." || dir == "" {
		return
	}
	_ = os.MkdirAll(dir, 0o755)
}

func (f *File) pathFor(level string) (string, bool) {
	level = normalizeLevelKey(level)
	if path, ok := f.files[level]; ok {
		return path, true
	}
	if path, ok := f.files["default"]; ok {
		return path, true
	}
	return "", false
}

func (f *File) loggerFor(path string) *log.Logger {
	f.mu.Lock()
	defer f.mu.Unlock()
	if lg, ok := f.loggers[path]; ok {
		return lg
	}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		panic("error opening file: " + path + ": " + err.Error())
	}
	lg := log.New(file, "", log.Ldate|log.Ltime)
	f.loggers[path] = lg
	return lg
}

func (f *File) Write(p []byte) (n int, err error) {
	if f.logger != nil {
		f.logger.Println(string(p))
		return len(p), nil
	}
	return len(p), nil
}

func (f *File) WriteWithLevel(p []byte, level string) (n int, err error) {
	if f.logger != nil {
		if !f.cfg.Exact {
			f.logger.Println(string(p))
		} else if f.level == level {
			f.logger.Println(string(p))
		}
		return len(p), nil
	}

	path, ok := f.pathFor(level)
	if !ok {
		return len(p), nil
	}
	f.loggerFor(path).Println(string(p))
	return len(p), nil
}
