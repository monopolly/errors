package errors

import (
	"fmt"

	"github.com/monopolly/jsons"
)

const (
	FieldID     = "id"    // string
	FieldCode   = "code"  // int
	FieldError  = "error" // string
	FieldPoints = "points"
)

// New error
func New(code int, id string, c ...any) (res E) {
	res = []byte("{}")
	res.Set(FieldCode, code)
	res.Set(FieldID, id)
	if c != nil {
		res.Set(FieldError, fmt.Sprint(c...))
	}
	return
}

func (a E) String() string {
	return string(a)
}

// error
type E []byte

// interface
func (a E) Error() string {
	return string(a)
}

// Set value
func (a *E) Set(k string, v any) {
	(*a) = jsons.Set((*a), k, v)
}

// Get value
func (a *E) Get(k string) jsons.Result {
	return jsons.Get((*a), k)
}

// Code set or get value
func (a *E) Code(v ...int) (res int) {
	if v == nil {
		return jsons.Int((*a), FieldCode)
	}
	a.Set(FieldCode, v[0])
	return
}

// Comment set or get value
func (a *E) Comment(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), FieldError)
	}
	a.Set(FieldError, v[0])
	return
}

// Append break points []string
func (a *E) AddPoint(v ...string) {
	switch len(v) {
	case 0:
		(*a) = jsons.ArrayStringAppend((*a), FieldPoints, BreakpointLevel(2))
	default:
		(*a) = jsons.ArrayStringAppend((*a), FieldPoints, v[0])
	}
}

// Comment set or get value
func (a *E) ID(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), FieldID)
	}
	a.Set(FieldID, v[0])
	return
}

func (a *E) Map() (res map[string]any) {
	return jsons.MapInterface(*a)
}
