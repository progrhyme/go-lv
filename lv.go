// Package lv implements a simple logger with filtering level.
//
// This package extends the standard package log and aimed to be minimal.
//
// Usage
//
// Using package-level logger:
//
//     import "github.com/progrhyme/go-lv"
//
//     func main() {
//         name := "Bob"
//         // lv.SetLevel(lv.LDebug) // for debug
//         lv.Infof("Hello, %s!", name)
//     }
//
// Creating a logger:
//
//     import (
//         "log"
//
//         "github.com/progrhyme/go-lv"
//     )
//
//     func main() {
//         logger := lv.New(os.Stderr, lv.LWarn, log.LstdFlags)
//         logger.Warnf("Something is wrong!")
//     }
package lv

// Minimal defines interface of minimal leveled logger.
type Minimal interface {
	GetLevel() Level
	SetLevel(Level)
	Print(...interface{})
	Printf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}

// Standard defines interface of standard leveled logger.
type Standard interface {
	Minimal
	Debug(...interface{})
	Debugf(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
}

// Granular defines interface of full leveled logger.
type Granular interface {
	Standard
	Trace(...interface{})
	Tracef(string, ...interface{})
	Notice(...interface{})
	Noticef(string, ...interface{})
	Critical(...interface{})
	Criticalf(string, ...interface{})
}
