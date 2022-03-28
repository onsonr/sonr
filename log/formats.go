package log

// Some predefined formats for logging output
const (
	// Convenient to use with client loggers, where we don't need fine-grained info
	// GetConnection ▶ INFO  Getting the connection
	FORMAT_SHORT = `%{color}%{shortfunc} ▶ %{level} %{color:reset} %{message}`

	// Same as FORMAT_SHORT but without the color information.
	// Appropriate for logging to files or on mobile clients.
	FORMAT_SHORT_COLORLESS = `%{shortfunc} ▶ %{level} %{message}`

	// Convenient to use with server loggers, where we need more fine-grained info and readable
	// output (includes color information, useful for console output)
	// [server][Mon 25.Sep 2017,14:11:041] Start ▶ NOTI  emmy server listening for connections on port 7007
	FORMAT_LONG = `%{color}[%{module}][%{time:Mon _2.Jan 2006,15:04:005}] %{shortfunc} ▶ %{level} %{color:reset} %{message}`

	// Same as FORMAT_LONG but without the color information (for files)
	FORMAT_LONG_COLORLESS = `[%{time:Mon _2.Jan 2006,15:04:005}] %{shortfunc} ▶ %{level} %{message}`
)
