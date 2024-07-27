package main

import (
	"bapt/cmd"
	"bapt/pkg/logger"
	"fmt"
)

func main() {
	// Initialize the logger
	logger.Initialize("bapt.log")

	if err := cmd.Execute(); err != nil {
		logger.ErrorLogger.Println(err)
		fmt.Println(err)
	}
}
