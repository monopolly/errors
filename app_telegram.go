package errors

import (
	"net/url"
	"strings"
)

// pgx error converter
func Telegram(err error) (e E) {

	if urlErr, ok := err.(*url.Error); ok {
		if urlErr.Timeout() {
			return Timeout(err)
		}
	}

	p := strings.ToLower(err.Error())
	switch {
	case strings.Contains(p, "forbidden"):
		return Forbidden(err)
	case strings.Contains(p, "too many requests"): //Too Many Requests: retry after 11 group:-4161715189
		return TooManyRequests(err)
	case strings.Contains(p, "can't parse"): // Bad Request: can't parse entities: Can't find end of the entity starting at byte offset 167
		return Unmarshal(err)
	default:
		return Unknown(err)
	}

}
