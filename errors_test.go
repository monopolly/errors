package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestRef(t *testing.T) {
	r := Access("comment1", "comment 2")
	r.SetRef()
	r.SetCodeLine()
	r.Message("a.me")
	fmt.Println(string(r))

	r.SetRef()
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
		New(401, "access", "some")
	}
}
