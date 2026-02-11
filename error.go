package errors

import jsoniter "github.com/json-iterator/go"

type Error struct {
	ID        string `json:"id,omitempty"`      // string
	UID       string `json:"uid,omitempty"`     // string
	Code      int    `json:"code,omitempty"`    // int
	App       string `json:"app,omitempty"`     // string
	Comment   string `json:"c,omitempty"`       // string
	Time      int64  `json:"time,omitempty"`    // int64
	Ref       string `json:"ref,omitempty"`     // string
	Trace     string `json:"trace,omitempty"`   // string
	IP        string `json:"ip,omitempty"`      // string
	Useragent string `json:"ua,omitempty"`      // string
	OS        string `json:"os,omitempty"`      // string
	Device    string `json:"device,omitempty"`  // string
	Explain   string `json:"explain,omitempty"` // string
	Message   string `json:"message,omitempty"` // string
	Fix       string `json:"fix,omitempty"`     // string
	Level     int    `json:"level,omitempty"`   // int
	Point     string `json:"point,omitempty"`   // string
	Package   string `json:"pkg,omitempty"`     // string
	Function  string `json:"func,omitempty"`    // string
}

func (a *Error) Pack() (res []byte) {
	res, _ = jsoniter.Marshal(a)
	return
}
