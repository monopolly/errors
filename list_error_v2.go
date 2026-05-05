package errors

type E2 map[string]any

func New2(id string) E2 {
	return E2{"id": id}
}

func (a EE) E(comment ...any) (res *E2) {
	e := New2(a.String())
	e["code"] = a.Int()
	e["c"] = comment
	return
}
