package errors

import (
	"net/url"
)

// pgx error converter
func ErrorURL(err error) (res *url.Error) {
	res, _ = err.(*url.Error)
	return
}
