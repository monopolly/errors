package errors

type List []E

func (a *List) Add(err E, limit int) {
	if len((*a)) > limit {
		(*a) = (*a)[:limit]
	}

	(*a) = append([]E{err}, (*a)...)
}

func (a *List) Reset() {
	(*a) = []E{}
}

func (a *List) Last() (err E) {
	if len((*a)) == 0 {
		return
	}

	return (*a)[0]
}
