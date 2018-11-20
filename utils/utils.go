package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var linkRex = regexp.MustCompile("[^a-zA-Z0-9]+")

func MD5String(value string) string {
	h := md5.New()
	io.WriteString(h, value)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func StandardiseLink(v string) string {
	return linkRex.ReplaceAllString(v, "")
}

func Files(folder string, opts ...string) (files []string) {
	ext := ".csv"
	if len(opts) > 0 {
		if x := opts[0]; x == "*" || x == "" {
			ext = ""
		} else {
			ext = x
		}
	}

	f, err := os.Open(folder)
	if err != nil {
		return
	}
	_ = f.Close()

	err = filepath.Walk(folder, func(path string, f os.FileInfo, err error) (e error) {
		if err != nil {
			return err
		}
		if !f.IsDir() && (ext == "" || strings.HasSuffix(path, ext)) {
			files = append(files, path)
		}
		return
	})

	return
}

func SetupFolders(folders ...string) error {
	fx := map[string]bool{}
	for _, f := range folders {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			fx[f] = true
		}
	}

	for f := range fx {
		if err := os.MkdirAll(f, 0700); err != nil {
			return err
		}

		if err := os.Chmod(f, 0700); err != nil {
			return err
		}
	}

	return nil
}

func FileExists(fname string) bool {
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		return false
	}

	return true
}

func Reopen(fname string) (*os.File, error) {
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		return os.Create(fname)
	}

	return os.OpenFile(fname, os.O_RDWR|os.O_APPEND, 0660)
}

func OpenLogfile(file string) (*os.File, error) {
	if err := SetupFolders(filepath.Dir(file)); err != nil {
		return nil, err
	}

	return Reopen(file)
}
