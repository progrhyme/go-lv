# go-lv

Package lv implements a simple logging package with filtering level. It
delegates output functionality to log.Logger from standard 'log' package.

# Usage

Using package Logger:

```go
import "github.com/progrhyme/go-lv"

func main() {
	name := "Bob"
	// lv.SetLevel(lv.Debug) // for debug
	lv.Infof("Hello, %s!", name)
}
```

Creating a Logger:

```go
import (
	"log"

	"github.com/progrhyme/go-lv"
)

func main() {
	logger := lv.New(os.Stderr, lv.Warning, log.LstdFlags)
	logger.Warnf("Something is wrong!")
}
```

# License

The MIT License.

Copyright (c) 2020 IKEDA Kiyoshi.
