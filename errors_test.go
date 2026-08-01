package errors

import (
	"fmt"
	"testing"
)

func TestBreakpoint(t *testing.T) {

	var p E //Unavailable("some")
	p.AddPoint()
	p.AddPoint("nice")
	p.Set("ok1", 44131)
	p.Set("ok2", true)
	p.Set("ok3", "yeee")

	fmt.Println(p.String())
}
