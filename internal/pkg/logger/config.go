package logger

type LoggingLevel string

const (
	DebugLevel LoggingLevel = "debug"
	ErrorLevel LoggingLevel = "error"
	InfoLevel  LoggingLevel = "info"
	WarnLevel  LoggingLevel = "warn"
	FatalLevel LoggingLevel = "fatal"
)

type Config struct {
	AppName string

	Level string `envconfig:"level"`
	Debug bool   `envconfig:"debug"`
}
