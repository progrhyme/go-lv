package lv

import (
	"fmt"
	"io"
	"log"
)

// Logger embeds standard log.Logger.
// Expanding it with filtering feature.
type Logger struct {
	*log.Logger
	level Level
}

// New creates a new Logger. The out variable sets the destination to which log data will be written.
// The prop argument defines the logging properties; it corresponds to 'flags' in 'log' package.
func New(out io.Writer, l Level, prop int) *Logger {
	return &Logger{Logger: log.New(out, "", prop), level: l}
}

// GetLevel returns filtering level of self.
func (self *Logger) GetLevel() (l Level) {
	return self.level
}

// SetLevel sets filtering level to self.
func (self *Logger) SetLevel(l Level) {
	self.level = l
}

func (self *Logger) Printf(format string, v ...interface{}) {
	self.Logger.Printf(format, v...)
}

func (self *Logger) writef(l Level, format string, v ...interface{}) {
	if l >= self.level {
		format := fmt.Sprintf("[%s] %s", l, format)
		self.Logger.Printf(format, v...)
	}
}

// Trace level logging.
func (self *Logger) Tracef(format string, v ...interface{}) {
	self.writef(LTrace, format, v...)
}

// Debug level logging.
func (self *Logger) Debugf(format string, v ...interface{}) {
	self.writef(LDebug, format, v...)
}

// Info level logging.
func (self *Logger) Infof(format string, v ...interface{}) {
	self.writef(LInfo, format, v...)
}

// Notice level logging.
func (self *Logger) Noticef(format string, v ...interface{}) {
	self.writef(LNotice, format, v...)
}

// Warning level logging.
func (self *Logger) Warnf(format string, v ...interface{}) {
	self.writef(LWarn, format, v...)
}

// Error level logging.
func (self *Logger) Errorf(format string, v ...interface{}) {
	self.writef(LError, format, v...)
}

// Critical level logging.
func (self *Logger) Criticalf(format string, v ...interface{}) {
	self.writef(LCritical, format, v...)
}

// Fatalf calls log.Fatalf() in standard package log.
func (self *Logger) Fatalf(format string, v ...interface{}) {
	self.Logger.Fatalf(format, v...)
}
