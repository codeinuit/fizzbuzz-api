package logger

type LogLevel string

const (
	WarnLevel  string = "warn"
	ErrorLevel string = "error"
	InfoLevel  string = "info"
	DebugLevel string = "debug"
)

// Logger interface implementation
type Logger interface {
	Info(...any)
	Warn(...any)
	Error(...any)
	Debug(...any)
}
