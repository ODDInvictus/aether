package logger

import (
	"strconv"
	"time"

	"github.com/fatih/color"
)

var debug bool

func Debug (b bool) {
	Log("Setting debug level to " + strconv.FormatBool(b))
	debug = b
}

func Warn(stmt string) {
	if !debug {
		return
	}

	color.Yellow("[%s] [Wrn] %s\n", time.Now().Format(time.RFC3339), stmt)
}

func Err(stmt string, err error) {
	color.Red("[%s] [Err] %s\n", time.Now().Format(time.RFC3339), stmt)
	panic(err)
}

func Log(stmt string) {
	color.Blue("[%s] [Log] %s\n", time.Now().Format(time.RFC3339), stmt)
}

func Verbose(stmt string) {
	if !debug {
		return
	}

	color.Magenta("[%s] [Ver] %s\n", time.Now().Format(time.RFC3339), stmt)
}