package lv

import (
	"fmt"
	"io"
	"os"
)

const defaultLevel Level = LInfo

var defaultLogger *Logger

// Configure sets given properties for package-level logger.
func Configure(out io.Writer, l Level, prop int) {
	defaultLogger = New(out, l, prop)
}

// GetLevel returns filtering level of package-level logger.
func GetLevel() Level {
	return getDefaultLogger().GetLevel()
}

// SetLevel sets filtering level for package-level logger.
func SetLevel(l Level) {
	getDefaultLogger().SetLevel(l)
}

func getDefaultLogger() *Logger {
	if defaultLogger == nil {
		Configure(os.Stderr, defaultLevel, 0)
	}
	return defaultLogger
}

// Print always prints given data like fmt.Print() with no prefix using package-level logger.
func Print(v ...interface{}) {
	getDefaultLogger().Logger.Output(2, fmt.Sprint(v...))
}

// Printf always prints given data like fmt.Printf() with no prefix using package-level logger.
func Printf(format string, v ...interface{}) {
	getDefaultLogger().Logger.Output(2, fmt.Sprintf(format, v...))
}

// Trace level logging using package-level logger.
func Trace(v ...interface{}) {
	getDefaultLogger().write(LTrace, v...)
}

// Tracef is another trace level logging using package-level logger.
func Tracef(format string, v ...interface{}) {
	getDefaultLogger().writef(LTrace, format, v...)
}

// Debug level logging using package-level logger.
func Debug(v ...interface{}) {
	getDefaultLogger().write(LDebug, v...)
}

// Debugf is another debug level logging using package-level logger.
func Debugf(format string, v ...interface{}) {
	getDefaultLogger().writef(LDebug, format, v...)
}

// Info level logging using package-level logger.
func Info(v ...interface{}) {
	getDefaultLogger().write(LInfo, v...)
}

// Infof is another info level logging using package-level logger.
func Infof(format string, v ...interface{}) {
	getDefaultLogger().writef(LInfo, format, v...)
}

// Notice level logging using package-level logger.
func Notice(v ...interface{}) {
	getDefaultLogger().write(LNotice, v...)
}

// Noticef is another notice level logging using package-level logger.
func Noticef(format string, v ...interface{}) {
	getDefaultLogger().writef(LNotice, format, v...)
}

// Warn is warning level logging using package-level logger.
func Warn(v ...interface{}) {
	getDefaultLogger().write(LWarn, v...)
}

// Warnf is another warning level logging using package-level logger.
func Warnf(format string, v ...interface{}) {
	getDefaultLogger().writef(LWarn, format, v...)
}

// Error level logging using package-level logger.
func Error(v ...interface{}) {
	getDefaultLogger().write(LError, v...)
}

// Errorf is another error level logging using package-level logger.
func Errorf(format string, v ...interface{}) {
	getDefaultLogger().writef(LError, format, v...)
}

// Critical level logging using package-level logger.
func Critical(v ...interface{}) {
	getDefaultLogger().write(LCritical, v...)
}

// Criticalf is another critical level logging using package-level logger.
func Criticalf(format string, v ...interface{}) {
	getDefaultLogger().writef(LCritical, format, v...)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	getDefaultLogger().Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	getDefaultLogger().Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}
