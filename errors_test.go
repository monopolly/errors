package errors

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestRef(t *testing.T) {
	v := Access("comment")
	v.Ref("L")
	v.Trace(time.Now().Unix(), "nice")
	v.Trace(time.Now().Unix(), "127.0.0.1")
	fmt.Println(string(v))

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
