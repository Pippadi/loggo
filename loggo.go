package loggo

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Level int

const (
	// DebugLevel is the lowest log level
	DebugLevel Level = iota

	// InfoLevel is the default log level, and is higher than
	// DebugLevel but lower than ErrorLevel
	InfoLevel

	// ErrorLevel is the highest log level
	ErrorLevel
)

const (
	redClr     = "\033[31m"
	orangeClr  = "\033[33m"
	greenClr   = "\033[92m"
	defaultClr = "\033[39m"
)

var prefixes = [3]string{
	orangeClr + "DBG" + defaultClr,
	greenClr + "INF" + defaultClr,
	redClr + "ERR" + defaultClr,
}

var level Level = InfoLevel
var outfile = os.Stdout

// SetLevel sets log level to one of DebugLevel, InfoLevel, or ErrorLevel.
// Log priority is given by DebugLevel < InfoLevel < ErrorLevel.
// Log level is InfoLevel by default.
func SetLevel(lvl Level) {
	level = lvl
}

// SetLevelString sets log level based on the string passed, which must be
// one of "debug", "info", and "error". Returns true if so, false otherwise.
func SetLevelString(lvl string) bool {
	var levelMap = map[string]Level{
		"debug": DebugLevel,
		"info":  InfoLevel,
		"error": ErrorLevel,
	}

	aLevel, ok := levelMap[strings.ToLower(lvl)]
	if ok {
		level = aLevel
	}
	return ok
}

// SetFile sets a log file. It is the calling function's responsibility to
// open and close the file. Default output is os.Stdout.
func SetFile(f *os.File) {
	outfile = f
}

// Debug prints arguments followed by a newline at DebugLevel.
func Debug(args ...any) {
	printRaw(DebugLevel, args...)
}

// Debugf prints arguments formatted by the given template at DebugLevel.
// A newline character is always appended to the final string.
func Debugf(template string, args ...any) {
	Debug(fmt.Sprintf(template, args...))
}

// Info prints arguments followed by a newline at InfoLevel.
func Info(args ...any) {
	printRaw(InfoLevel, args...)
}

// Infof prints arguments formatted by the given template at InfoLevel.
// A newline character is always appended to the final string.
func Infof(template string, args ...any) {
	Info(fmt.Sprintf(template, args...))
}

// Error prints arguments followed by a newline at ErrorLevel.
func Error(args ...any) {
	printRaw(ErrorLevel, args...)
}

// Errorf prints arguments formatted by the given template at ErrorLevel.
// A newline character is always appended to the final string.
func Errorf(template string, args ...any) {
	Error(fmt.Sprintf(template, args...))
}

// Trace is meant to be used with Un, to record the entering and leaving of fuctions
// Usage: defer loggo.Un(loggo.Trace("MyFunc"))
func Trace(fName string) string {
	Debug("Entering", fName)
	return fName
}

// Un is meant to be used with Trace, to record the entering and leaving of fuctions
// Usage: defer loggo.Un(loggo.Trace("MyFunc"))
func Un(fName string) {
	Debug("Leaving", fName)
}

// printRaw conditionally prints args with the current
// timestamp and colorized log level prefix
func printRaw(lvl Level, args ...any) {
	if lvl >= level {
		outfile.WriteString(fmt.Sprintln(append(
			[]any{time.Now().Local().Format("2006-01-02T15:04:05.000") + " [" + prefixes[lvl] + "]"},
			args...)...,
		))
	}
}
