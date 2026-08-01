package errors

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func Breakpoint(short ...bool) (res string) {
	_, file, line, _ := runtime.Caller(1)
	switch len(short) {
	case 0:
		return fmt.Sprintf("%s:%d", file, line)
	default:
		_, file = filepath.Split(file)
		return fmt.Sprintf("%s:%d", file, line)
	}
}

func BreakpointLevel(level int, short ...bool) (res string) {
	_, file, line, _ := runtime.Caller(level)
	switch len(short) {
	case 0:
		return fmt.Sprintf("%s:%d", file, line)
	default:
		_, file = filepath.Split(file)
		return fmt.Sprintf("%s:%d", file, line)
	}
}
