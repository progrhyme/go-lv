package lv

import (
	"strings"
)

// Level represents filtering level.
// Logger will filter out log data when called with lower level functions
type Level int

const (
	LTrace Level = iota + 1
	LDebug
	LInfo // Default level for package-level logger
	LNotice
	LWarning
	LError
)

var labelToLevel map[string]Level

// String returns human-readable string of log level.
//
// Examples:
//
//     LDebug.String()  //=> "DEBUG"
//     LNotice.String() //=> "NOTICE"
func (lv Level) String() string {
	switch lv {
	case LTrace:
		return "TRACE"
	case LDebug:
		return "DEBUG"
	case LInfo:
		return "INFO"
	case LNotice:
		return "NOTICE"
	case LWarning:
		return "WARNING"
	case LError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// LabelToLevel returns Level by given label.
//
// Examples:
//
//     lv.LabelToLevel("trace")     //=> lv.LTrace
//     lv.LabelToLevel("warning")   //=> lv.LWarning
//     lv.LabelToLevel("undefined") //=> 0
func LabelToLevel(label string) Level {
	if labelToLevel == nil {
		initLabelToLevel()
	}
	return labelToLevel[label]
}

func initLabelToLevel() {
	labelToLevel = make(map[string]Level)
	for i := LTrace; i.String() != "UNKNOWN"; i++ {
		labelToLevel[strings.ToLower(i.String())] = i
	}
}
