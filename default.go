package lv

import (
	"io"
	"log"
	"os"
	"strings"
)

const defaultLevel Level = Info

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
func Printf(fmt string, v ...interface{}) {
	getDefaultLogger().Printf(fmt, v...)
}

// Trace level logging using package-level logger.
func Tracef(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[TRACE]", fmt}, " ")
	getDefaultLogger().writef(Trace, fmt, v...)
}

// Debug level logging using package-level logger.
func Debugf(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[DEBUG]", fmt}, " ")
	getDefaultLogger().writef(Debug, fmt, v...)
}

// Info level logging using package-level logger.
func Infof(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[INFO]", fmt}, " ")
	getDefaultLogger().writef(Info, fmt, v...)
}

// Notice level logging using package-level logger.
func Noticef(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[NOTICE]", fmt}, " ")
	getDefaultLogger().writef(Notice, fmt, v...)
}

// Warn level logging using package-level logger.
func Warnf(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[WARN]", fmt}, " ")
	getDefaultLogger().writef(Warning, fmt, v...)
}

// Error level logging using package-level logger.
func Errorf(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[ERROR]", fmt}, " ")
	getDefaultLogger().writef(Error, fmt, v...)
}

// Fatalf calls log.Fatalf() in standard package log through package-level logger.
func Fatalf(fmt string, v ...interface{}) {
	getDefaultLogger().Fatalf(fmt, v...)
}
