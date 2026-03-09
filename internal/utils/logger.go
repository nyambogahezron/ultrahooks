package utils

import (
	"fmt"
	"os"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
)

func Info(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}

func Success(msg string, args ...interface{}) {
	fmt.Printf(ColorGreen+"✔ "+msg+ColorReset+"\n", args...)
}

func Warning(msg string, args ...interface{}) {
	fmt.Printf(ColorYellow+"⚠ "+msg+ColorReset+"\n", args...)
}

func Error(msg string, args ...interface{}) {
	fmt.Printf(ColorRed+"✖ "+msg+ColorReset+"\n", args...)
}

func Header(msg string) {
	fmt.Printf(ColorYellow+"\n%s"+ColorReset+"\n\n", msg)
}

func ProcessName(msg string) {
	fmt.Printf("\n[%s]\n", msg)
}

func CommandLog(msg string) {
	fmt.Printf("%s\n", msg)
}

func Fatal(msg string, args ...interface{}) {
	Error(msg, args...)
	os.Exit(1)
}
