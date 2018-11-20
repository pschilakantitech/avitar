package pidfile

import (
	"fmt"
	"os"
	"path/filepath"
)

var pidFile string

// Dump writes the pid to disk.
func Dump() error {
	f, err := os.Create(pidFile)
	if err != nil {
		return err
	}

	fmt.Fprint(f, os.Getpid())
	return f.Close()
}

// Drop removes the pidfile.
func Drop() error {
	return os.Remove(pidFile)
}

func init() {
	pidFile = filepath.Base(os.Args[0]) + ".pid"
}
