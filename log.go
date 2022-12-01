package main

import (
	"fmt"
	"log"
	"os"
)

var errLogger *log.Logger

func init() {
	pre := fmt.Sprintf("%s: %s", appName, "[ERR]")
	errLogger = log.New(os.Stderr, pre, log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
}

func Err(err error) {
	errLogger.Println(err.Error())
}
