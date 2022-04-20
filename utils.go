package logger

import "github.com/go-zoox/datetime"

// Now returns the current time string.
func Now() string {
	return datetime.Now().Format("YYYY-MM-DD HH:mm:ss")
}
