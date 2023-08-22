package utils

import (
	"os"
	"runtime"
)

// IsFile returns true if given path exists as a file (i.e. not a directory).
func IsFile(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsDir returns true if given path is a directory, and returns false when it's
// a file or does not exist.
func IsDir(dir string) bool {
	f, e := os.Stat(dir)
	if e != nil {
		return false
	}
	return f.IsDir()
}

// IsExist returns true if a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// CurrentUsername returns the current system user via environment variables.
func CurrentUsername() string {
	username := os.Getenv("USER")
	if len(username) > 0 {
		return username
	}

	return os.Getenv("USERNAME")
}

// CurrentOSVer returns the current system version via golang runtime
func CurrentOSVer() string {
	return runtime.GOOS
}

// CurrentArch returns the current os arch via golang runtime
func CurrentArch() string {
	return runtime.GOARCH
}
