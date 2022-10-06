package logger

import (
	"errors"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is the main logger that is abstracted in this package.
// It is exported here for use with gorm.
var Logger *logrus.Logger

// ErrInvalidLevel is returned when an invalid log level is given in the config
var ErrInvalidLevel = errors.New("invalid log level")

// Config represents configuration details for logging.
type Config struct {
	Filename string `json:"filename"`
	Level    string `json:"level"`
}

func init() {
	Logger = logrus.New()
	Logger.Formatter = &logrus.TextFormatter{DisableColors: true}
}

// Setup configures the logger based on options in the config.json.
func Setup(config *Config) error {
	var err error
	// Set up logging level
	level := logrus.InfoLevel
	if config.Level != "" {
		level, err = logrus.ParseLevel(config.Level)
		if err != nil {
			return err
		}
	}
	Logger.SetLevel(level)
	// Set up logging to a file if specified in the config
	logFile := config.Filename
	if logFile != "" {
		f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		mw := io.MultiWriter(os.Stderr, f)
		Logger.Out = mw
	}
	return nil
}

// Debug logs a debug message
func Debug(args ...any) {
	Logger.Debug(args...)
}

// Debugf logs a formatted debug messsage
func Debugf(format string, args ...any) {
	Logger.Debugf(format, args...)
}

// Info logs an informational message
func Info(args ...any) {
	Logger.Info(args...)
}

// Infof logs a formatted informational message
func Infof(format string, args ...any) {
	Logger.Infof(format, args...)
}

// Error logs an error message
func Error(args ...any) {
	Logger.Error(args...)
}

// Errorf logs a formatted error message
func Errorf(format string, args ...any) {
	Logger.Errorf(format, args...)
}

// Warn logs a warning message
func Warn(args ...any) {
	Logger.Warn(args...)
}

// Warnf logs a formatted warning message
func Warnf(format string, args ...any) {
	Logger.Warnf(format, args...)
}

// Fatal logs a fatal error message
func Fatal(args ...any) {
	Logger.Fatal(args...)
}

// Fatalf logs a formatted fatal error message
func Fatalf(format string, args ...any) {
	Logger.Fatalf(format, args...)
}

// WithFields returns a new log enty with the provided fields
func WithFields(fields logrus.Fields) *logrus.Entry {
	return Logger.WithFields(fields)
}

// Writer returns the current logging writer
func Writer() *io.PipeWriter {
	return Logger.Writer()
}
