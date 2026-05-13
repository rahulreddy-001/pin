package main

import (
	"os"

	"github.com/rahulreddy-001/pin"
)

func main() {
	loggerConfig := pin.NewLoggerConfig().WithName("main-application-1").WithWriter(os.Stdout).WithLogLevel(pin.LEVEL_STAGE)
	logger := pin.NewLogger(loggerConfig)
	defer logger.Flush()

	logger.Info("sample info from main", pin.Any("file", "main.go"), pin.Any("cli", "true"))
}
