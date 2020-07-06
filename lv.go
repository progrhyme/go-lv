// Package lv implements a simple logging package with filtering level.
// It delegates output functionality to log.Logger from standard 'log' package.
package lv

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
