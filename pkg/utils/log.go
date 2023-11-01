package utils

import (
	"fmt"
	"os"
	"time"
)

func Log(level, message string) {
	now := time.Now()
	fmt.Printf("%s %s: %s", now.Format("15:04:05"), level, message)
}

func Panic(err error) {
	if err != nil {
		Log("ERROR", fmt.Sprintf("%s\n", err.Error()))
		os.Exit(1)
	}
}
