package constants

const LogOutputBuffer = 1024

const (
	// LevelDebug is Level Debug
	LevelDebug = "DEBUG"
	// LevelInfo is Level Info
	LevelInfo = "INFO"
	// LevelWarn is Level Warn
	LevelWarn = "WARN"
	// LevelError is Level Error
	LevelError = "ERROR"
	// LevelFatal is Level Fatal
	LevelFatal = "FATAL"
)

var LevelMap = map[string]int{
	"DEBUG": 0,
	"INFO":  1,
	"WARN":  2,
	"ERROR": 3,
	"FATAL": 4,
}
