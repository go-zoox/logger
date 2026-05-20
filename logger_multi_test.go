package logger

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	cst "github.com/go-zoox/logger/components/transport"
	"github.com/go-zoox/logger/transport/console"
	"github.com/go-zoox/logger/transport/file"
)

func TestLogger_consoleAndByLevelFile(t *testing.T) {
	dir := t.TempDir()
	errPath := filepath.Join(dir, "error.log")

	transports := map[string]cst.Transport{
		"console": console.New(),
		"file": file.New(&file.Config{
			Levels: map[string]string{"error": errPath},
		}),
	}

	l := New(func(opt *Option) {
		opt.Transports = transports
		opt.Level = "DEBUG"
	})

	l.Info("info-msg")
	l.Error("error-msg")

	eb, err := os.ReadFile(errPath)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(eb), "error-msg") {
		t.Fatalf("error file: %q", eb)
	}
}
