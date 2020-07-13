[![latest-tag](https://badgen.net/github/tag/progrhyme/go-lv)](https://github.com/progrhyme/go-lv/releases)
[![go-test](https://github.com/progrhyme/go-lv/workflows/go-test/badge.svg)](https://github.com/progrhyme/go-lv/actions?query=workflow%3Ago-test)

# go-lv

Package lv implements a simple logger with filtering level.

This package extends the standard package log and aimed to be minimal.

# Usage

Using package-level logger:

```go
import "github.com/progrhyme/go-lv"

func main() {
	name := "Bob"
	// lv.SetLevel(lv.LDebug) // for debug
	lv.Infof("Hello, %s!", name)
}
```

Creating a logger:

```go
import (
	"log"

	"github.com/progrhyme/go-lv"
)

func main() {
	logger := lv.New(os.Stderr, lv.LWarn, log.LstdFlags)
	logger.Warn("Something is wrong!")
}
```

# Documenatation

https://pkg.go.dev/github.com/progrhyme/go-lv

# Limitation

No support for structured logging.

# License

The MIT License.

Copyright (c) 2020 IKEDA Kiyoshi.
