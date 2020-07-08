package lv

import (
	"io"
	"log"
	"os"
)

const defaultLevel Level = LInfo

var defaultLogger *Logger

// Configure sets given properties for package-level logger.
func Configure(out io.Writer, lv Level, prop int) {
	defaultLogger = &Logger{Logger: log.New(out, "", prop), Level: lv}
}

// GetLevel returns filtering level of package-level logger.
func GetLevel() Level {
	return getDefaultLogger().Level
}

// SetLevel sets filtering level for package-level logger.
func SetLevel(lv Level) {
	getDefaultLogger().Level = lv
}

func getDefaultLogger() *Logger {
	if defaultLogger == nil {
		Configure(os.Stderr, defaultLevel, 0)
	}
	return defaultLogger
}

// Printf always prints given data like fmt.Printf() with no prefix using package-level logger.
func Printf(format string, v ...interface{}) {
	getDefaultLogger().Printf(format, v...)
}

// Trace level logging using package-level logger.
func Tracef(format string, v ...interface{}) {
	getDefaultLogger().writef(LTrace, format, v...)
}

// Debug level logging using package-level logger.
func Debugf(format string, v ...interface{}) {
	getDefaultLogger().writef(LDebug, format, v...)
}

// Info level logging using package-level logger.
func Infof(format string, v ...interface{}) {
	getDefaultLogger().writef(LInfo, format, v...)
}

// Notice level logging using package-level logger.
func Noticef(format string, v ...interface{}) {
	getDefaultLogger().writef(LNotice, format, v...)
}

// Warning level logging using package-level logger.
func Warnf(format string, v ...interface{}) {
	getDefaultLogger().writef(LWarning, format, v...)
}

// Error level logging using package-level logger.
func Errorf(format string, v ...interface{}) {
	getDefaultLogger().writef(LError, format, v...)
}

// Fatalf calls log.Fatalf() in standard package log through package-level logger.
func Fatalf(format string, v ...interface{}) {
	getDefaultLogger().Fatalf(format, v...)
}
