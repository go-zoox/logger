package file

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNew_singleFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "app.log")
	tr := New(&Config{FilePath: path})
	tr.WriteWithLevel([]byte("line\n"), "INFO")
	body, _ := os.ReadFile(path)
	if !strings.Contains(string(body), "line") {
		t.Fatalf("got %q", body)
	}
}

func TestNew_levelsWithDefaultFile(t *testing.T) {
	dir := t.TempDir()
	errPath := filepath.Join(dir, "error.log")
	defPath := filepath.Join(dir, "app.log")

	tr := New(&Config{
		FilePath: defPath,
		Levels: map[string]string{
			"error": errPath,
		},
	})
	tr.WriteWithLevel([]byte("err-line\n"), "ERROR")
	tr.WriteWithLevel([]byte("info-line\n"), "INFO")

	eb, _ := os.ReadFile(errPath)
	if !strings.Contains(string(eb), "err-line") {
		t.Fatalf("error log: %q", eb)
	}
	db, _ := os.ReadFile(defPath)
	if !strings.Contains(string(db), "info-line") {
		t.Fatalf("default log: %q", db)
	}
}

func TestNew_levelsOnlyUsesFilePathAsDefault(t *testing.T) {
	dir := t.TempDir()
	defPath := filepath.Join(dir, "only.log")

	tr := New(&Config{FilePath: defPath, Levels: map[string]string{}})
	// Levels empty => single-file mode
	tr.WriteWithLevel([]byte("x\n"), "WARN")
	body, _ := os.ReadFile(defPath)
	if !strings.Contains(string(body), "x") {
		t.Fatalf("got %q", body)
	}
}
