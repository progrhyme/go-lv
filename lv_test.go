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
	if !doLog(New(out, Info, 0)) {
		t.Errorf("New Logger is not Loggable")
	}
}

func TestFilteredLogging(t *testing.T) {
	out := &strings.Builder{}
	Configure(out, Info, 0)
	msg := "blah blah blah."
	Debugf(msg)
	if out.String() != "" {
		t.Errorf("Log is not filtered out. Got: %s, Want: %s", out.String(), "")
	}
	Warnf(msg)
	want := fmt.Sprintf("[WARN] %s\n", msg)
	if out.String() != want {
		t.Errorf("Log is printed. Got: %s, Want: %s", out.String(), want)
	}
}

func TestLabelToLevel(t *testing.T) {
	label2Lv := map[string]Level{
		"trace": Trace, "warning": Warning, "undef": 0,
	}
	for label, lv := range label2Lv {
		if LabelToLevel(label) != lv {
			t.Errorf("Label unmatch. Got %s, Want %s", LabelToLevel(label), lv)
		}
	}
}
