package log

// Setup initializes logging to a default logger.
// For the most use cases/when you only want to log to a SINGLE logfile using this
// library, this should do. Otherwise use New().
func Setup(cfg Config) (err error) {
	defaultLogger, err = New(cfg, cfg.Writer)

	return
}

// Debug logs information tagged as [DEBUG] and only if debugging is enabled.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Info logs information tagged as [INFO]
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Warn logs information tagged as [WARN]
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Error logs information tagged as [ERROR]
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Fatal logs information tagged as [FATAL], prints the message to the screen and aborts.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}
