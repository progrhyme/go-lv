package lv

import (
	"io"
	"log"
	"os"
	"strings"
)

const defaultLevel Level = Info

var defaultLogger *Logger

// Configure sets given properties for package Logger.
func Configure(out io.Writer, lv Level, prop int) {
	defaultLogger = &Logger{Logger: log.New(out, "", prop), Level: lv}
}

// GetLevel returns filtering level of package Logger.
func GetLevel() Level {
	return getDefaultLogger().Level
}

// SetLevel sets filtering level for package Logger.
func SetLevel(lv Level) {
	getDefaultLogger().Level = lv
}

func getDefaultLogger() *Logger {
	if defaultLogger == nil {
		Configure(os.Stderr, defaultLevel, 0)
	}
	return defaultLogger
}

// Printf always prints given data like fmt.Printf() with no prefix using package Logger.
func Printf(fmt string, v ...interface{}) {
	getDefaultLogger().Printf(fmt, v...)
}

// Trace level logging using package Logger.
func Tracef(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[TRACE]", fmt}, " ")
	getDefaultLogger().writef(Trace, fmt, v...)
}

// Debug level logging using package Logger.
func Debugf(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[DEBUG]", fmt}, " ")
	getDefaultLogger().writef(Debug, fmt, v...)
}

// Info level logging using package Logger.
func Infof(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[INFO]", fmt}, " ")
	getDefaultLogger().writef(Info, fmt, v...)
}

// Notice level logging using package Logger.
func Noticef(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[NOTICE]", fmt}, " ")
	getDefaultLogger().writef(Notice, fmt, v...)
}

// Warn level logging using package Logger.
func Warnf(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[WARN]", fmt}, " ")
	getDefaultLogger().writef(Warning, fmt, v...)
}

// Error level logging using package Logger.
func Errorf(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[ERROR]", fmt}, " ")
	getDefaultLogger().writef(Error, fmt, v...)
}

// Calls 'Fatalf' of standard log package through package Logger.
func Fatalf(fmt string, v ...interface{}) {
	getDefaultLogger().Fatalf(fmt, v...)
}
