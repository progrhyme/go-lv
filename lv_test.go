package lv

import (
	"fmt"
	"strings"
	"testing"
)

func TestLoggable(t *testing.T) {
	doLog := func(lg Granular) bool {
		return true
	}
	if !doLog(defaultLogger) {
		t.Errorf("defaultLogger is not granular logger")
	}
	out := &strings.Builder{}
	if !doLog(New(out, LInfo, 0)) {
		t.Errorf("New Logger is not granular logger")
	}
}

func TestFilteredLogging(t *testing.T) {
	msg := "blah blah blah."
	level := LNotice
	tests := []struct {
		name   string
		output string
		log    func(s string)
	}{
		{"Infof", "", func(s string) { Infof(s) }},
		{"Notice", fmt.Sprintf("[NOTICE] %s\n", msg), func(s string) { Notice(s) }},
		{"Warnf", fmt.Sprintf("[WARN] %s\n", msg), func(s string) { Warnf(s) }},
		{"Error", fmt.Sprintf("[ERROR] %s\n", msg), func(s string) { Error(s) }},
		{"Criticalf", fmt.Sprintf("[CRITICAL] %s\n", msg), func(s string) { Criticalf(s) }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &strings.Builder{}
			Configure(out, level, 0)
			tt.log(msg)
			if out.String() != tt.output {
				if tt.output == "" {
					t.Errorf("Log is not filtered out. Got: %s, Want: %s", out.String(), tt.output)
				} else {
					t.Errorf("Output log differs. Got: %s, Want: %s", out.String(), tt.output)
				}
			}
		})
	}
}

func TestWordToLevel(t *testing.T) {
	label2Lv := map[string]Level{
		"trace": LTrace, "Info": LInfo, "WARN": LWarn, "critical": LCritical, "undef": 0,
	}
	for label, lv := range label2Lv {
		if WordToLevel(label) != lv {
			t.Errorf("Label unmatch. Got %s, Want %s", WordToLevel(label), lv)
		}
	}
}
