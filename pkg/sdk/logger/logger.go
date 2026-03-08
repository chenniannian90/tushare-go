package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *logrus.Logger

// LogConfig defines the configuration for the logger
type LogConfig struct {
	Filename   string // Log file path, empty means stdout
	MaxSize    int    // Maximum size in megabytes
	MaxAge     int    // Maximum number of days to retain old log files
	MaxBackups int    // Maximum number of old log files to retain
	Compress   bool   // Compress rotated files
	Level      string // Log level: debug, info, warn, error
	Format     string // Log format: json, text
}

// Init initializes the global logger with the given configuration
func Init(cfg *LogConfig) {
	log = logrus.New()

	// Set output based on configuration
	if cfg.Filename != "" {
		// Use lumberjack for log rotation
		log.SetOutput(&lumberjack.Logger{
			Filename:   cfg.Filename,
			MaxSize:    cfg.MaxSize,    // MB
			MaxAge:     cfg.MaxAge,     // days
			MaxBackups: cfg.MaxBackups, // number of backups
			Compress:   cfg.Compress,   // compress rotated files
		})
	} else {
		log.SetOutput(os.Stdout)
	}

	// Set log level
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	log.SetLevel(level)

	// Set formatter based on format config
	if cfg.Format == "json" {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}
}

// SetOutput sets the logger output (useful for testing)
func SetOutput(w io.Writer) {
	if log != nil {
		log.SetOutput(w)
	}
}

// L returns the global logger instance
func L() *logrus.Logger {
	if log == nil {
		// Fallback to default logger if not initialized
		log = logrus.New()
		log.SetOutput(os.Stdout)
		log.SetLevel(logrus.InfoLevel)
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}
	return log
}

// WithField creates an entry with a single field
func WithField(key string, value interface{}) *logrus.Entry {
	return L().WithField(key, value)
}

// WithFields creates an entry with multiple fields
func WithFields(fields logrus.Fields) *logrus.Entry {
	return L().WithFields(fields)
}

// Fields is an alias for logrus.Fields
type Fields = logrus.Fields

// WithError creates an entry with an error field
func WithError(err error) *logrus.Entry {
	return L().WithError(err)
}

// Debug logs a message at level Debug
func Debug(args ...interface{}) {
	L().Debug(args...)
}

// Debugf logs a formatted message at level Debug
func Debugf(format string, args ...interface{}) {
	L().Debugf(format, args...)
}

// Info logs a message at level Info
func Info(args ...interface{}) {
	L().Info(args...)
}

// Infof logs a formatted message at level Info
func Infof(format string, args ...interface{}) {
	L().Infof(format, args...)
}

// Warn logs a message at level Warn
func Warn(args ...interface{}) {
	L().Warn(args...)
}

// Warnf logs a formatted message at level Warn
func Warnf(format string, args ...interface{}) {
	L().Warnf(format, args...)
}

// Error logs a message at level Error
func Error(args ...interface{}) {
	L().Error(args...)
}

// Errorf logs a formatted message at level Error
func Errorf(format string, args ...interface{}) {
	L().Errorf(format, args...)
}

// Fatal logs a message at level Fatal then the process will exit with status set to 1
func Fatal(args ...interface{}) {
	L().Fatal(args...)
}

// Fatalf logs a formatted message at level Fatal then the process will exit with status set to 1
func Fatalf(format string, args ...interface{}) {
	L().Fatalf(format, args...)
}
