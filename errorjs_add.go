package errors

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"path/filepath"
	"runtime"
	"time"

	"github.com/monopolly/jsons"
	"github.com/rs/xid"
)

var (
	// lower = "0123456789qwertyuiopasdfghjklzxcvbnm"
	// upper = "0123456789QWERTYUIOPASDFGHJKLZXCVBNM"
	// hh, _ = hashids.NewWithData(&hashids.HashIDData{Alphabet: lower})
	guid = xid.New()
)

// New error
func New(code int, id string, comment ...interface{}) (a E) {

	switch len(comment) > 0 {
	case true:
		return []byte(fmt.Sprintf(`{"code":%d,"id":"%s","time":%d,"c":"%v","ref":"%s%d"}`, code, id, time.Now().Unix(), jsonEscape(fmt.Sprint(comment[0])), guid.String(), rand.Intn(1000)))
	case false:
		return []byte(fmt.Sprintf(`{"code":%d,"id":"%s","time":%d,"ref":"%s%d"}`, code, id, time.Now().Unix(), guid.String(), rand.Intn(1000)))
	}
	return
}

// github.com/monopolly/errors.TestRef errors_test.go:25
func (a *E) Point() {
	function, file, line, _ := runtime.Caller(1)
	_, file = filepath.Split(file)
	a.Set(eSource, fmt.Sprintf("%s %s:%d", runtime.FuncForPC(function).Name(), file, line))
}

// github.com/monopolly/errors.TestRef errors_test.go:25
func (a *E) Line() string {
	return jsons.String((*a), eSource)
}

// Ref set or get value
func (a *E) Ref(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), eRef)
	}
	a.Set(eRef, v[0])
	return
}

// Ref set or get value
func (a *E) Nil() (res bool) {
	return len(*a) == 0
}

// Go set go link
func (a *E) Go(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), "go")
	}
	a.Set("go", v[0])
	return
}

func (a *E) Base64() (res string) {
	return base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(*a)
}
func (a E) String() string {
	return string(a)
}

func ParseBase64(b string) (res E, err error) {
	return base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(b)
}
