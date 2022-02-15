package trail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAncestor(t *testing.T) {
	trail := testTrail
	a := trail.Ancestors()

	assert.Equal(t, "/Users/name/personal/trail", a.Next().Inner)
	assert.Equal(t, "/Users/name/personal", a.Next().Inner)
	assert.Equal(t, "/Users/name", a.Next().Inner)
	assert.Equal(t, "/Users", a.Next().Inner)
	assert.Nil(t, a.Next())
	assert.Nil(t, a.Next())
}
