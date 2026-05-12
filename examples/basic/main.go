package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rahulreddy-001/pin"
)

type CLILogEncoder struct{}

func (_ CLILogEncoder) Encode(log pin.Log) string {
	return fmt.Sprintf("[%s]  [%s] %s %s  %s  %s", log.Level, log.Name, log.Source, log.Timestamp.Local().Format(time.RFC1123), log.Message, log.Fields.Encode())
}

func main() {
	logger := pin.NewLogger("main-application-1", pin.LEVEL_STAGE, os.Stdout, CLILogEncoder{})
	defer logger.Flush()

	logger.Info("sample info from main", pin.Any("file", "main.go"), pin.Any("cli", "true"))
}
