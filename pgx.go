package errors

import (
	ee "errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// pgx error converter
func PGX(er error) (err E) {

	pc, filename, line, _ := runtime.Caller(1)
	defer err.Set("source", fmt.Sprintf("%s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err))
	e := er.Error()
	switch er {
	case pgx.ErrNoRows:
		return NotFound(er)
	}

	var p *pgconn.PgError
	if ee.As(er, &p) {
		switch p.Code {
		case pgerrcode.NoDataFound, pgerrcode.CaseNotFound:
			return NotFound(er)
		case pgerrcode.UniqueViolation, pgerrcode.DuplicateAlias, pgerrcode.DuplicateColumn, pgerrcode.DuplicateTable, pgerrcode.DuplicateJSONObjectKeyValue:
			return Exist(p.ConstraintName)
		default:
		}
	}

	switch {
	case strings.Contains(e, "23505"):
		return Exist(er)
	case strings.Contains(e, "rows is closed"):
		return NotFound(er)
	}

	return Database(er)

}
