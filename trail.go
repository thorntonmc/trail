package trail

import "errors"

type Trail struct {
	Inner string
}

func NewTrail(s string) Trail {
	return Trail{s}
}

func (t Trail) Components() (*Components, error) {
	hpr, err := hasPhysicalRoot(t.asByteSlice())
	if err != nil {
		return nil, err
	}

	return &Components{
		trail:           t.asByteSlice(),
		hasPhysicalRoot: hpr,
	}, nil
}

func (t Trail) asByteSlice() []byte {
	return []byte(t.Inner)
}

func hasPhysicalRoot(b []byte) (bool, error) {
	if len(b) <= 0 {
		return false, errors.New("empty []byte passed")
	}
	return isSepByte(b[0]), nil
}

func (t Trail) Ancestors() *Ancestor {
	return &Ancestor{&t}
}
