package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	config Config
	lg     *log.Logger
}

func New(conf Config) (*Logger, error) {
	if conf.Level != "" {
		setLevel(conf.Level)
	}

	if conf.Debug {
		setDebugLevel()
	}

	return &Logger{
		config: conf,
		lg:     log.New(os.Stderr, conf.AppName, log.LstdFlags),
	}, nil
}

func logLevel(level string) (l LoggingLevel) {
	switch level {
	case "debug":
		return DebugLevel
	case "error":
		return ErrorLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "fatal":
		return FatalLevel
	default:
		panic("unknown log level")
	}
}

func setLevel(level string) {
	os.Setenv("LOG_LEVEL", string(logLevel(level)))
}

func setDebugLevel() {
	os.Setenv("LOG_LEVEL", string(DebugLevel))
}

func (l Logger) LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.lg.Printf("[Error]: %s\n", msg)
}

func (l Logger) LogInfo(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.lg.Printf("[Info]: %s\n", msg)
}

func (l Logger) LogFatalf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.lg.Printf("[Fatal]: %s\n", msg)
}
