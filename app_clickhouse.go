package errors

import (
	"database/sql"
	ee "errors"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func ClickHouse(er error) (err E) {
	if er == nil {
		return nil
	}

	if ee.Is(er, sql.ErrNoRows) {
		return NotFound(er)
	}

	e := er.Error()

	var ex *clickhouse.Exception
	if ee.As(er, &ex) {
		switch ex.Code {
		case 60, 81, 47:
			return NotFound(er)

		case 57, 82:
			return Exist(er)

		default:
			return Database(er)
		}
	}

	switch {
	case strings.Contains(e, "no rows in result set"),
		strings.Contains(e, "rows is closed"),
		strings.Contains(e, "UNKNOWN_TABLE"),
		strings.Contains(e, "UNKNOWN_DATABASE"),
		strings.Contains(e, "UNKNOWN_IDENTIFIER"),
		strings.Contains(e, "doesn't exist"):
		return NotFound(er)

	case strings.Contains(e, "TABLE_ALREADY_EXISTS"),
		strings.Contains(e, "DATABASE_ALREADY_EXISTS"),
		strings.Contains(e, "already exists"):
		return Exist(er)

	case strings.Contains(e, "connection refused"),
		strings.Contains(e, "no such host"),
		strings.Contains(e, "broken pipe"),
		strings.Contains(e, "connection reset by peer"):
		return Connection(er)
	}

	return Database(er)
}
