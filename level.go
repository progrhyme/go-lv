package lv

import (
	"strings"
)

// Level represents filtering level.
// Logger will filter out log data when called with lower level functions
type Level int

const (
	Trace Level = iota + 1
	Debug
	Info // Default level for package-level logger
	Notice
	Warning
	Error
)

var labelToLevel map[string]Level

// String returns human-readable string of log level.
//
// Examples:
//
//     Debug.String()  //=> "Debug"
//     Notice.String() //=> "Notice"
func (lv Level) String() string {
	switch lv {
	case Trace:
		return "Trace"
	case Debug:
		return "Debug"
	case Info:
		return "Info"
	case Notice:
		return "Notice"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	default:
		return "Unknown"
	}
}

// LabelToLevel returns Level by given label.
//
// Examples:
//
//     lv.LabelToLevel("trace")     //=> lv.Trace
//     lv.LabelToLevel("warning")   //=> lv.Warning
//     lv.LabelToLevel("undefined") //=> 0
func LabelToLevel(label string) Level {
	if labelToLevel == nil {
		initLabelToLevel()
	}
	return labelToLevel[label]
}

func initLabelToLevel() {
	labelToLevel = make(map[string]Level)
	for i := Trace; i.String() != "Unknown"; i++ {
		labelToLevel[strings.ToLower(i.String())] = i
	}
}
