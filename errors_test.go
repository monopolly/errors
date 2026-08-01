package errors

import (
	"fmt"
	"testing"
)

func TestBreakpoint(t *testing.T) {

	p := Unavailable("some")
	p.AddPoint()

	fmt.Println(p.String())
}
