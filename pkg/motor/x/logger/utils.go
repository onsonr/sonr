package logger

func conevertLogLevel(level LogLevel) string {
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
