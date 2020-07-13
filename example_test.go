package lv

import (
	"log"
	"os"
)

func Example() {
	// Using package-level logger
	name := "Bob"
	// lv.SetLevel(lv.LDebug) // for debug
	Infof("Hello, %s!", name)

	// Creating a logger
	logger := New(os.Stderr, LWarn, log.LstdFlags)
	logger.Warn("Something is wrong!")
}

func ExampleConfigure() {
	Configure(os.Stderr, LNotice, log.LstdFlags|log.Lshortfile)
	Noticef("Version is undefined")
}
