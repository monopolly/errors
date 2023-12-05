package errors

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestTrace(t *testing.T) {

	fmt.Println(errors.Unwrap(Access()))

	fmt.Println("Trace")
	r := Access()
	r.Trace("one")
	r.Trace("two")
	r.Trace("three")

	fmt.Println(string(r))

	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond * 3)
		fmt.Println(Access())
	}

}

func TestRef(t *testing.T) {
	r := Access("comment1", "comment 2")
	r.Point()
	r.Message("a.me")
	fmt.Println(string(r))

}

func BenchmarkPreErr(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Backup()
	}
}

func BenchmarkGolangError(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		errors.New("some error")
	}
}

func BenchmarkNewError(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		New(401, "access", "some1")
	}
}
