package errors

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"path/filepath"
	"runtime"
	"time"

	"github.com/monopolly/jsons"
	"github.com/speps/go-hashids"
)

var (
	lower = "0123456789qwertyuiopasdfghjklzxcvbnm"
	upper = "0123456789QWERTYUIOPASDFGHJKLZXCVBNM"
	hh, _ = hashids.NewWithData(&hashids.HashIDData{Alphabet: lower})
	App   = "" //app add automatic if has
)

// New error
func New(code int, id string, comment ...interface{}) (a E) {
	app := fmt.Sprintf(`"app":"%s"`, App)
	switch len(comment) > 0 {
	case true:
		return []byte(fmt.Sprintf(`{%s,"code":%d,"id":"%s","time":%d,"c":"%v"}`, app, code, id, time.Now().Unix(), jsonEscape(fmt.Sprint(comment[0]))))
	case false:
		return []byte(fmt.Sprintf(`{%s,"code":%d,"id":"%s","time":%d}`, app, code, id, time.Now().Unix()))
	}

	return
}

// hashids[user, unix] || random
func (a *E) SetRef(userID ...int) {
	var h string
	switch userID {
	case nil:
		h, _ = hh.Encode([]int{
			int(time.Now().Unix()),
			rand.Intn(99),
			rand.Intn(99),
			rand.Intn(999),
		})
	default:
		h, _ = hh.Encode([]int{userID[0], int(time.Now().Unix())})
	}

	a.Ref(h)
}

// SetRef ID
func (a *E) R() {
	h, _ := hh.Encode([]int{int(time.Now().Unix()), rand.Intn(99), rand.Intn(99), rand.Intn(999)})
	a.Ref(h)
}

// Error system
func (a E) Error() string {
	return string(a)
}

// Error system
func (a E) Err() error {
	return fmt.Errorf(string(a))
}

// SetCodeLine set point in code
func (a *E) SetCodeLine() {
	function, file, line, _ := runtime.Caller(1)
	_, file = filepath.Split(file)
	a.Set(eSource, fmt.Sprintf("%s %s:%d", runtime.FuncForPC(function).Name(), file, line))
}

// CodeLine set or get value
func (a *E) CodeLine() (res string) {
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

func ParseBase64(b string) (res E, err error) {
	return base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(b)
}
