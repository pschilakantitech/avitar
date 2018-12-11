package log

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/chilakantip/avitar/utils"
)

// Logger is the proxy to all the logging methods.
type Logger struct {
	Config
	io.Writer
	*log.Logger
}

// Config groups the library settings.
type Config struct {
	LogName,
	LogPrefix,
	AppName,
	AppEnv string
	Debug bool

	Writer io.Writer
}

var defaultLogger *Logger

// New initializes a logger.
func New(cfg Config, w io.Writer) (l *Logger, err error) {
	l = &Logger{Config: cfg, Writer: w}

	if l.Writer == nil {
		var err error
		l.Writer, err = utils.OpenLogfile(l.LogName)
		if err != nil {
			return nil, err
		}
	}

	l.Logger = log.New(l.Writer, l.LogPrefix+" ", log.Ldate|log.Lmicroseconds)

	return
}

// Debug logs information tagged as [DEBUG] and only if debugging is enabled.
func (l *Logger) Debug(args ...interface{}) {
	if l.Config.Debug {
		l.logentry("[DEBUG]", args...)
	}
}

// Info logs information tagged as [INFO]
func (l *Logger) Info(args ...interface{}) {
	l.logentry("[INFO]", args...)
}

// Warn logs information tagged as [WARN]
func (l *Logger) Warn(args ...interface{}) {
	l.logentry("[WARN]", args...)
}

// Error logs information tagged as [ERROR]
func (l *Logger) Error(args ...interface{}) {
	l.logentry("[ERROR]", args...)
}

// Fatal logs information tagged as [FATAL], prints the message to the screen and aborts.
func (l *Logger) Fatal(args ...interface{}) {
	l.logentry("[FATAL]", args...)

	os.Exit(255)
}

// Close should be called when the log is no longer needed.
// NOTE: It is especially important for long running processes that may open many logfiles,
// to appropriately close the no longer needed logfiles (or risk running out of file descriptors).
func (l *Logger) Close() error {
	if f, ok := l.Writer.(*os.File); ok {
		return f.Close()
	} else if _, ok := l.Writer.(*bytes.Buffer); ok {
		// nothing to do
	} else {
		return fmt.Errorf("Don't know how to close the writer")
	}

	return nil
}

func (l *Logger) logentry(kind string, args ...interface{}) {
	all := []interface{}{kind}
	all = append(all, args...)
	if l != nil && l.Logger != nil {
		l.Println(all...)
	} else {
		fmt.Println("ERROR: Applog not initialized.")
		fmt.Println(all...)
	}
}
