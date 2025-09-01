package utils

import (
	"fmt"
	"log"
	"os"
)

func DebugLog(logMsg string) {
	f, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	logger := log.New(
		f,
		"DEBUG: ",
		log.LstdFlags,
	)
	logger.Println(logMsg)
}
