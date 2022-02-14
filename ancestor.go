package main

import "path"

// Path represents a slash("/") separated path
type Path string

func (p Path) String() string {
	return string(p)
}

func (p Path) Ancestors() *Ancestor {
	return &Ancestor{
		p,
	}
}

type Ancestor struct {
	val Path
}

func (a *Ancestor) Next() Path {
	p := a.val
	a.val = Path(path.Dir(a.val.String()))
	return p
}
