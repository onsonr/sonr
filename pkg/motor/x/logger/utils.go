package logger

func conevertLogLevelCategory(level LogLevel) string {
	switch level {
	case LEVEL_INFO:
		return "[INFO]"
	case LEVEL_DEBUG:
		return "[DEBUG]"
	case LEVEL_WARN:
		return "[WARN]"
	case LEVEL_ERROR:
		return "[ERROR]"
	case LEVEL_FATAL:
		return "[FATAL]"
	}

	return ""
}

func conevertLogLevel(level string) LogLevel {
	switch level {
	case "info":
		return LEVEL_INFO
	case "debug":
		return LEVEL_DEBUG
	case "warn":
		return LEVEL_WARN
	case "error":
		return LEVEL_ERROR
	case "fatal":
		return LEVEL_FATAL
	}

	return LEVEL_WARN
}
