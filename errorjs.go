package errors

/*
	Global:
	[] - tupple jsons [1,"name"]
	! - New вместо NewItem
	noinit - no New()

	Fields:
	time - Time() function for int
	nofunc - no field func
	++ add multiply func for int
	++ simple func for boolean
	= new json name ex: =money.count or =company.report.name
*/
import (
	"time"

	"github.com/monopolly/jsons"
)

const (
	eID        = "id"      // string
	eCode      = "code"    // int
	eApp       = "app"     // string
	eComment   = "c"       // string
	eTime      = "time"    // int64 time
	eSource    = "at"      // string
	eRef       = "ref"     // string
	eTrace     = "trace"   // string duration
	eIP        = "ip"      // string
	eUseragent = "ua"      // string user agent
	eOS        = "os"      // string
	eDevice    = "device"  // string
	eExplain   = "explain" // string
	eMessage   = "message" // string for public message: Sorry, server is not available
	eFix       = "fix"     // string how to fix it: restart server
	eUser      = "uid"     // int
	eLevel     = "level"   // int
)

// E is a struct
type E []byte

// error interface
func (a E) Error() string {
	return string(a)
}

// Set value
func (a *E) Set(k string, v interface{}) *E {
	(*a) = jsons.Set((*a), k, v)
	return a
}

// Get value
func (a *E) Get(k string) jsons.Result {
	return jsons.Get((*a), k)
}

// Get value
func (a *E) Bytes() []byte {
	return []byte((*a))
}

// ID set or get value
func (a *E) ID(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), eID)
	}
	a.Set(eID, v[0])
	return
}

// Code set or get value
func (a *E) Code(v ...int) (res int) {
	if v == nil {
		return jsons.Int((*a), eCode)
	}
	a.Set(eCode, v[0])
	return
}

// App set or get value
func (a *E) App(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), eApp)
	}
	a.Set(eApp, v[0])
	return
}

// App set or get value
func (a *E) Explain(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), eExplain)
	}
	a.Set(eExplain, v[0])
	return
}

// Comment set or get value
func (a *E) Comment(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), eComment)
	}
	a.Set(eComment, v[0])
	return
}

// Time set or get value
func (a *E) Time(v ...int64) (res int64) {
	if v == nil {
		return jsons.Int64((*a), eTime)
	}
	a.Set(eTime, v[0])
	return
}

// TimeTime get value as time
func (a *E) TimeTime() (res time.Time) {
	return time.Unix(a.Time(), 0)
}

// Trace set or get value
func (a *E) Trace(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), eTrace)
	}
	(*a) = jsons.ArrayStringAppend((*a), eTrace, v[0])
	return
}

// IP set or get value
func (a *E) IP(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), eIP)
	}
	a.Set(eIP, v[0])
	return
}

// Useragent set or get value
func (a *E) Useragent(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), eUseragent)
	}
	a.Set(eUseragent, v[0])
	return
}

// OS set or get value
func (a *E) OS(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), eOS)
	}
	a.Set(eOS, v[0])
	return
}

// Device set or get value
func (a *E) Device(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), eDevice)
	}
	a.Set(eDevice, v[0])
	return
}

// Message set public message
func (a *E) Message(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), eMessage)
	}
	a.Set(eMessage, v[0])
	return
}

// Fix set or get value
func (a *E) Fix(v ...string) (res string) {
	if v == nil {
		return jsons.String((*a), eFix)
	}
	a.Set(eFix, v[0])
	return
}

// User set or get value
func (a *E) User(v ...int) (res int) {
	if v == nil {
		return jsons.Int((*a), eUser)
	}
	a.Set(eUser, v[0])
	return
}

// User set or get value
func (a *E) Level(v ...int) (res int) {
	if v == nil {
		return jsons.Int((*a), eLevel)
	}
	a.Set(eLevel, v[0])
	return
}

func (a *E) Map() (res map[string]any) {
	return jsons.MapInterface(*a)
}
