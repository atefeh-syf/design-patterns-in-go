package main
import (
	"fmt"
)

type LegacyLogger struct {}

func (l *LegacyLogger) LogMessage(message string) {
    fmt.Println("Legacy logger: " + message)
}

type NewLogger interface {
    Log(message string)
}

type NewSystemLogger struct {}

func (l *NewSystemLogger) Log(message string) {
    fmt.Println("New system logger: " + message)
}

type LoggerAdapter struct {
    legacyLogger *LegacyLogger
}

func (a *LoggerAdapter) Log(message string) {
    a.legacyLogger.LogMessage(message)
}

func logMessage(message string, logger NewLogger) {
    logger.Log(message)
}

func main() {
    legacyLogger := &LegacyLogger{}
    loggerAdapter := &LoggerAdapter{legacyLogger: legacyLogger}

    logMessage("Hello, world!", loggerAdapter)
}

