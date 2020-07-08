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

// Loggable determines interface of Logger
type Loggable interface {
	Printf(string, ...interface{})
	Tracef(string, ...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Noticef(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
}
