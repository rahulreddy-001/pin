package pin

type LogLevel int

var (
	LEVEL_STAGE LogLevel = -1
	LEVEL_PROD  LogLevel = 0
	DEBUG       LogLevel = 0
	INFO        LogLevel = 1
	WARN        LogLevel = 2
	ERROR       LogLevel = 3
	FATAL       LogLevel = 4
)

func (this LogLevel) IsGreaterThan(level LogLevel) bool {
	return this > level
}

func (this LogLevel) GetLogLevel() string {
	switch this {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "INFO"
	}
}
