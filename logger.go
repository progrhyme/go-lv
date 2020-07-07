package lv

import (
	"io"
	"log"
	"strings"
)

// Logger embeds standard log.Logger.
// Expanding it with filtering feature.
type Logger struct {
	*log.Logger
	Level Level // Filtering level. Lowers will be filtered out
}

// New creates a new Logger. The out variable sets the destination to which log data will be written.
// The prop argument defines the logging properties; it corresponds to 'flags' in 'log' package.
func New(out io.Writer, lv Level, prop int) *Logger {
	return &Logger{Logger: log.New(out, "", prop), Level: lv}
}

func (self *Logger) Printf(fmt string, v ...interface{}) {
	self.Logger.Printf(fmt, v...)
}

func (self *Logger) writef(lv Level, fmt string, v ...interface{}) {
	if lv >= self.Level {
		self.Logger.Printf(fmt, v...)
	}
}

// Trace level logging.
func (self *Logger) Tracef(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[TRACE]", fmt}, " ")
	self.writef(Trace, fmt, v...)
}

// Debug level logging.
func (self *Logger) Debugf(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[DEBUG]", fmt}, " ")
	self.writef(Debug, fmt, v...)
}

// Info level logging.
func (self *Logger) Infof(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[INFO]", fmt}, " ")
	self.writef(Info, fmt, v...)
}

// Notice level logging.
func (self *Logger) Noticef(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[NOTICE]", fmt}, " ")
	self.writef(Notice, fmt, v...)
}

// Warning level logging.
func (self *Logger) Warnf(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[WARN]", fmt}, " ")
	self.writef(Warning, fmt, v...)
}

// Error level logging.
func (self *Logger) Errorf(fmt string, v ...interface{}) {
	fmt = strings.Join([]string{"[ERROR]", fmt}, " ")
	self.writef(Error, fmt, v...)
}

// Fatalf calls log.Fatalf() in standard package log.
func (self *Logger) Fatalf(fmt string, v ...interface{}) {
	self.Logger.Fatalf(fmt, v...)
}
