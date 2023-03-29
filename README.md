# loggo

An absurdly simple go logging package supporting three log levels

# Example

```go
package main

import (
	"os"

	log "github.com/Pippadi/loggo"
)

func main() {
	logfile, err := os.Create("test.log")
	if err != nil {
		log.Errorf("Unable to read log file! Reason: %s", err.Error())
	}
	defer logfile.Close()
	log.SetFile(logfile)

	// Default log level is Info, so this won't show up
	log.Debug("Starting up")

	log.SetLevel(log.DebugLevel)

	// Now it will
	log.Debug("Set correct logging level")

	n := 4
	log.Info("Changing the world", n, "lines at a time")

	log.Error("Barf!")
}
```

Outputs

```
2023-03-29T14:08:33.751 [DBG]  Set correct logging level
2023-03-29T14:08:33.751 [INF]  Changing the world 4 lines at a time
2023-03-29T14:08:33.751 [ERR]  Barf!
```
