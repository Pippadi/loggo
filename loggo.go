package loggo

import (
	"fmt"
	"os"
	"time"
)

// Requirements:
// - Minimal interruption to calling function
// - Redirection to file (optional)

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	ErrorLevel
)

const (
	redClr     = "\033[31m"
	orangeClr  = "\033[33m"
	defaultClr = "\033[39m"
)

var prefixes = [3]string{
	orangeClr + "DBG" + defaultClr,
	"INF",
	redClr + "ERR" + defaultClr,
}

var level Level = InfoLevel
var outfile = os.Stdout

func SetLevel(lvl Level) {
	level = lvl
}
func SetFile(f *os.File) {
	outfile = f
}

func Debug(msgs ...any) {
	printRaw(DebugLevel, msgs...)
}
func Debugf(template string, args ...any) {
	Debug(fmt.Sprintf(template, args...))
}

func Info(msgs ...any) {
	printRaw(InfoLevel, msgs...)
}
func Infof(template string, args ...any) {
	Info(fmt.Sprintf(template, args...))
}

func Error(msgs ...any) {
	printRaw(ErrorLevel, msgs...)
}
func Errorf(template string, args ...any) {
	Error(fmt.Sprintf(template, args...))
}

func printRaw(lvl Level, args ...any) {
	if lvl >= level {
		outfile.WriteString(fmt.Sprintln(append(
			[]any{time.Now().Local().Format("2006-01-02T15:04:05.000") + " [" + prefixes[lvl] + "] "},
			args...)...,
		))
	}
}
