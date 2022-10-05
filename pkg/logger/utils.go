package logger

func conevertLogLevelCategory(level LogLevel) string {
	switch level {
	case LEVEL_INFO:
		return CAT_INFO
	case LEVEL_DEBUG:
		return CAT_DEBUG
	case LEVEL_WARN:
		return CAT_WARN
	case LEVEL_ERROR:
		return CAT_ERROR
	case LEVEL_FATAL:
		return CAT_FATAL
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
