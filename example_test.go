package lv_test

import (
	"log"
	"os"

	"github.com/progrhyme/go-lv"
)

func Example() {
	// Using package-level logger
	name := "Bob"
	// lv.SetLevel(lv.LDebug) // for debug
	lv.Infof("Hello, %s!", name)

	// Creating a logger
	logger := lv.New(os.Stderr, lv.LWarn, log.LstdFlags)
	logger.Warn("Something is wrong!")
}

func ExampleConfigure() {
	lv.Configure(os.Stderr, lv.LNotice, log.LstdFlags|log.Lshortfile)
	lv.Noticef("Version is undefined")
}
