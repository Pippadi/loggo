package loggo

import (
	"errors"
	"math/rand"
	"os"
	"testing"
)

func logSamples() {
	Debug("Hi")
	Infof("%d is a number", 234)
	Error("Barf", errors.New("Sample error"))
}

func TestDebug(t *testing.T) {
	SetLevel(DebugLevel)
	logSamples()
}

func TestInfo(t *testing.T) {
	SetLevel(InfoLevel)
	logSamples()
}

func TestError(t *testing.T) {
	SetLevel(ErrorLevel)
	logSamples()
}

func TestLogToFile(t *testing.T) {
	logfile, err := os.Create("/tmp/loggo.log")
	if err != nil {
		return
	}
	defer logfile.Close()
	SetLevel(DebugLevel)
	SetFile(logfile)
	for i := 0; i < 1000; i++ {
		lvl := rand.Int31()
		if lvl > 2000000000 {
			Info("This is Info")
		} else {
			Debug("This is Debug")
		}
	}
}

func TestUnTrace(t *testing.T) {
	SetFile(os.Stdout)
	defer Un(Trace("TestUnTrace"))
}
