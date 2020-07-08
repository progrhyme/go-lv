package lv

import "strings"

// Level represents filtering level.
// Logger will filter out log data when called with lower level functions
type Level int

const (
	LTrace Level = iota + 1
	LDebug
	LInfo // Default level for package-level logger
	LNotice
	LWarn
	LError
	LCritical
)

var strToLevel map[string]Level

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
	case LWarn:
		return "WARN"
	case LError:
		return "ERROR"
	case LCritical:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

// WordToLevel returns Level by given word. Argument word is case-insensitive.
//
// Examples:
//
//     lv.WordToLevel("trace")     //=> lv.LTrace
//     lv.WordToLevel("Info")      //=> lv.LInfo
//     lv.WordToLevel("WARN")      //=> lv.LWarn
//     lv.WordToLevel("undefined") //=> 0
func WordToLevel(word string) Level {
	if strToLevel == nil {
		initStringToLevel()
	}
	return strToLevel[strings.ToUpper(word)]
}

func initStringToLevel() {
	strToLevel = make(map[string]Level)
	for i := LTrace; i.String() != "UNKNOWN"; i++ {
		strToLevel[i.String()] = i
	}
}
