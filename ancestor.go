package trail

import "path"

type Ancestor struct {
	trail *Trail
}

func (a *Ancestor) Next() *Trail {
	base := path.Base(string(a.trail.Inner))
	if a.trail.Inner == base {
		return nil
	}

	trail := a.trail
	nt := NewTrail(path.Dir(a.trail.Inner))
	a.trail = &nt
	return trail
}
