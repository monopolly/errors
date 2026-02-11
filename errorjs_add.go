package errors

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/monopolly/jsons"
	"github.com/sqids/sqids-go"
)

var (
	// lower = "0123456789qwertyuiopasdfghjklzxcvbnm"
	// upper = "0123456789QWERTYUIOPASDFGHJKLZXCVBNM"
	// hh, _ =  hashids.NewWithData(&hashids.HashIDData{Alphabet: lower})
	// guid  = xid.New()

	sids, _ = sqids.New(sqids.Options{Alphabet: "0123456789qwertyuiopasdfghjklzxcvbnm"})
	ref     = func(code int) (res string) {
		res, _ = sids.Encode([]uint64{uint64(code), uint64(time.Now().UnixMicro())})
		return
	}
)

// New error
func New(code int, id string, c ...any) (res E) {
	p := Error{
		Code: code,
		ID:   id,
		Ref:  ref(code), //.fmt.Sprintf("%d%d", code, time.Now().UnixMicro()), // uuid.NewString(),
	}
	if c != nil {
		p.Comment = fmt.Sprint(c...)
	}
	res = p.Pack()
	res.BreakpointLight(3)
	return
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

func (a *E) TelegramString() (res string) {
	var list []string
	list = append(list, fmt.Sprint(a.Code()), ":", a.ID())
	list = append(list, a.Comment())
	return strings.Join(list, "\n")
}

func ParseBase64(b string) (res E, err error) {
	return base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(b)
}
