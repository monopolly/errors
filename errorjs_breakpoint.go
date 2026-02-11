package errors

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

const ()

func (a *E) Point() string    { return a.Get(ePoint).String() }
func (a *E) Package() string  { return a.Get(ePackage).String() }
func (a *E) Function() string { return a.Get(eFunction).String() }

func (a *E) Breakpoint(level int) {
	if level == 0 {
		level = 3
	}
	f, file, line, _ := runtime.Caller(level)
	_, file = filepath.Split(file)
	a.Set(ePoint, fmt.Sprintf("%s:%d", file, line))
	a.Set(ePackage, trimFuncName(runtime.FuncForPC(f).Name()))
	a.Set(eFunction, getFuncName(runtime.FuncForPC(f).Name()))
}

func (a *E) BreakpointLight(level int) {
	if level == 0 {
		level = 3
	}
	_, file, line, _ := runtime.Caller(level)
	_, file = filepath.Split(file)
	a.Set(ePoint, fmt.Sprintf("%s:%d", file, line))
	// a.Set(ePackage, trimFuncName(runtime.FuncForPC(f).Name()))
	// a.Set(eFunction, getFuncName(runtime.FuncForPC(f).Name()))
}

func trimFuncName(full string) string {
	if i := strings.LastIndex(full, "."); i != -1 {
		return full[:i]
	}
	return full
}

func getFuncName(full string) string {
	if i := strings.LastIndex(full, "."); i != -1 {
		return full[i+1:]
	}
	return full
}
