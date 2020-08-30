package errors

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/monopolly/jsons"
)

const (
	fid      = "id"
	fcode    = "code"
	fcomment = "comment"
	ftime    = "time"
	fsource  = "source"
	fref     = "ref"
	ftrace   = "trace"
)

func New(code int, id string, comment ...interface{}) (a E) {
	switch len(comment) > 0 {
	case true:
		return []byte(fmt.Sprintf(`{"code":%d,"id":"%s","time":%d,"comment":"%v"}`, code, id, time.Now().Unix(), comment[0]))
	case false:
		return []byte(fmt.Sprintf(`{"code":%d,"id":"%s","time":%d}`, code, id, time.Now().Unix()))
	}
	return
}

type E []byte

func (a *E) Set(k string, v interface{}) *E {
	(*a) = jsons.Set((*a), k, v)
	return a
}

func (a *E) Time() int64 {
	return jsons.Int64((*a), ftime)
}

func (a *E) ID() string {
	return jsons.String((*a), fid)
}

func (a *E) Code() int {
	return jsons.Int((*a), fcode)
}

func (a *E) Comment() string {
	return jsons.String((*a), fcomment)
}

func (a *E) Source() string {
	return jsons.String((*a), fsource)
}

func (a *E) Refs() string {
	return jsons.String((*a), fref)
}

func (a *E) Ref(prefix ...string) *E {
	r := time.Now().UnixNano() / 1000
	switch len(prefix) > 0 {
	case true:
		return a.Set(fref, fmt.Sprintf("%s%d", prefix[0], r))
	case false:
		return a.Set(fref, fmt.Sprint(r))
	}
	return a
}

/* type E struct {
	ID      string `json:"id,omitempty" msg:"id,omitempty"`
	Comment string `json:"comment,omitempty" msg:"comment,omitempty"`
	Time    int64  `json:"time,omitempty" msg:"time,omitempty"` //nano
	Code    int    `json:"code,omitempty" msg:"code,omitempty"`
	Source  string `json:"source,omitempty" msg:"source,omitempty"`
	Ref     string `json:"refs,omitempty" msg:"refs,omitempty"`
} */

func (e E) Error() string {
	return fmt.Sprintf("%d %s %s %s", e.Code(), e.ID(), e.Comment(), e.Source())
}

//сохраняет номер строки кода откуда вызвана ошибка
func (e *E) X() *E {
	function, file, line, _ := runtime.Caller(1)
	_, file = filepath.Split(file)
	e.Set(fsource, fmt.Sprintf("%s %s:%d", runtime.FuncForPC(function).Name(), file, line))
	return e
}

func (e *E) TimeFormat(format string) string {
	return time.Unix(e.Time(), 0).Format(format)
}

func (e *E) GetTime() time.Time {
	return time.Unix(e.Time(), 0)
}

func (e *E) SetCode(code int) *E {
	return e.Set(fcode, code)
}

//когда нужно с разных серверов
func (a *E) Trace(v ...interface{}) *E {
	(*a) = jsons.Set((*a), ftrace, append(jsons.ArrayString((*a), ftrace), fmt.Sprintln(v...)))
	return a
}

func itos(x interface{}) string {
	switch x.(type) {
	case int:
		return strconv.Itoa(x.(int))
	case int8:
		return strconv.Itoa(int(x.(int8)))
	case int16:
		return strconv.Itoa(int(x.(int16)))
	case int32:
		return strconv.Itoa(int(x.(int32)))
	case int64:
		return strconv.Itoa(int(x.(int64)))
	case uint:
		return strconv.Itoa(int(x.(uint)))
	case uint8:
		return strconv.Itoa(int(x.(uint8)))
	case uint16:
		return strconv.Itoa(int(x.(uint16)))
	case uint32:
		return strconv.Itoa(int(x.(uint32)))
	case uint64:
		return strconv.Itoa(int(x.(uint64)))
	case string:
		return x.(string)
	case []byte:
		return string(x.([]byte))
	case bool:
		switch x.(bool) {
		case true:
			return "true"
		default:
			return "false"
		}
	case error:
		return x.(error).Error()
	default:
		return fmt.Sprintf("%v", x)
	}
}
