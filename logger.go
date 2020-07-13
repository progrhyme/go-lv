package lv

import (
	"fmt"
	"io"
	"log"
	"os"
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

// Print calls log.Logger.Output to print log. Arguments are handled in the manner of fmt.Print.
func (self *Logger) Print(v ...interface{}) {
	self.Logger.Output(2, fmt.Sprint(v...))
}

// Printf calls log.Logger.Output to print log. Arguments are handled in the manner of fmt.Printf.
func (self *Logger) Printf(format string, v ...interface{}) {
	self.Logger.Output(2, fmt.Sprintf(format, v...))
}

func (self *Logger) write(l Level, v ...interface{}) {
	if l >= self.level {
		v = append([]interface{}{fmt.Sprintf("[%s] ", l)}, v...)
		self.Logger.Output(3, fmt.Sprint(v...))
	}
}

func (self *Logger) writef(l Level, format string, v ...interface{}) {
	if l >= self.level {
		format := fmt.Sprintf("[%s] %s", l, format)
		self.Logger.Output(3, fmt.Sprintf(format, v...))
	}
}

// Trace level logging.
func (self *Logger) Trace(v ...interface{}) {
	self.write(LTrace, v...)
}

// Tracef is another trace level logging with format.
func (self *Logger) Tracef(format string, v ...interface{}) {
	self.writef(LTrace, format, v...)
}

// Debug level logging.
func (self *Logger) Debug(v ...interface{}) {
	self.write(LDebug, v...)
}

// Debugf is another debug level logging with format.
func (self *Logger) Debugf(format string, v ...interface{}) {
	self.writef(LDebug, format, v...)
}

// Info level logging.
func (self *Logger) Info(v ...interface{}) {
	self.write(LInfo, v...)
}

// Infof is another info level logging with format.
func (self *Logger) Infof(format string, v ...interface{}) {
	self.writef(LInfo, format, v...)
}

// Notice level logging.
func (self *Logger) Notice(v ...interface{}) {
	self.write(LNotice, v...)
}

// Noticef is another notice level logging with format.
func (self *Logger) Noticef(format string, v ...interface{}) {
	self.writef(LNotice, format, v...)
}

// Warn is warning level logging.
func (self *Logger) Warn(v ...interface{}) {
	self.write(LWarn, v...)
}

// Warnf is another warning level logging with format.
func (self *Logger) Warnf(format string, v ...interface{}) {
	self.writef(LWarn, format, v...)
}

// Error level logging.
func (self *Logger) Error(v ...interface{}) {
	self.write(LError, v...)
}

// Errorf is another error level logging with format.
func (self *Logger) Errorf(format string, v ...interface{}) {
	self.writef(LError, format, v...)
}

// Critical level logging.
func (self *Logger) Critical(v ...interface{}) {
	self.write(LCritical, v...)
}

// Criticalf is another critical level logging with format.
func (self *Logger) Criticalf(format string, v ...interface{}) {
	self.writef(LCritical, format, v...)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func (self *Logger) Fatal(v ...interface{}) {
	self.Logger.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func (self *Logger) Fatalf(format string, v ...interface{}) {
	self.Logger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}
