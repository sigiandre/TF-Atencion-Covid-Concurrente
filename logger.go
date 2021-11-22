package main

import (
	"log"
	"os"
	"strings"

	"github.com/go-errors/errors"
)

func Log(level string, msg string, msgOpt ...interface{}) {
	var toLog []interface{}
	toLog = append(toLog, "["+strings.ToUpper(level)+"]", msg)
	toLog = append(toLog, msgOpt...)
	log.Println(toLog...)
}

func Info(msg string, msgOpt ...interface{}) {
	Log("info", msg, msgOpt...)
}

func Warning(msg string, msgOpt ...interface{}) {
	Log("warn", msg, msgOpt...)
}

func Error(msg string, msgOpt ...interface{}) {
	Log("error", msg, msgOpt...)
}

func Stack(err error) {
	Log("error-STACK", errors.Wrap(err, 1).ErrorStack())
}

func FatalPanic(msg string, msgOpt ...interface{}) {

	Log("error-panic", msg, msgOpt...)
	panic(1)
}

func FatalExit(msg string, msgOpt ...interface{}) {
	Log("error-exit", msg, msgOpt...)
	os.Exit(1)
}
