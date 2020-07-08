package lv

import (
	"fmt"
	"strings"
	"testing"
)

func TestLoggable(t *testing.T) {
	doLog := func(lg Loggable) bool {
		return true
	}
	if !doLog(defaultLogger) {
		t.Errorf("defaultLogger is not Loggable")
	}
	out := &strings.Builder{}
	if !doLog(New(out, LInfo, 0)) {
		t.Errorf("New Logger is not Loggable")
	}
}

func TestFilteredLogging(t *testing.T) {
	out := &strings.Builder{}
	Configure(out, LNotice, 0)
	msg := "blah blah blah."
	Infof(msg)
	if out.String() != "" {
		t.Errorf("Log is not filtered out. Got: %s, Want: %s", out.String(), "")
	}
	Warnf(msg)
	want := fmt.Sprintf("[WARNING] %s\n", msg)
	if out.String() != want {
		t.Errorf("Log is printed. Got: %s, Want: %s", out.String(), want)
	}
}

func TestWordToLevel(t *testing.T) {
	label2Lv := map[string]Level{
		"trace": LTrace, "Info": LInfo, "WARNING": LWarning, "undef": 0,
	}
	for label, lv := range label2Lv {
		if WordToLevel(label) != lv {
			t.Errorf("Label unmatch. Got %s, Want %s", WordToLevel(label), lv)
		}
	}
}
