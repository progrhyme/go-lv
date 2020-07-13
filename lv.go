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
	Printf(string, ...interface{})
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
}

// Standard defines interface of standard leveled logger.
type Standard interface {
	Minimal
	Debugf(string, ...interface{})
	Warnf(string, ...interface{})
}

// Granular defines interface of full leveled logger.
type Granular interface {
	Standard
	Tracef(string, ...interface{})
	Noticef(string, ...interface{})
	Criticalf(string, ...interface{})
}
