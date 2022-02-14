package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAncestor(t *testing.T) {
	p := Path("/Users/michaelthornton/personal/ancestor")
	a := p.Ancestors()

	assert.Equal(t, "/Users/michaelthornton/personal/ancestor", a.Next().String())
	assert.Equal(t, "/Users/michaelthornton/personal", a.Next().String())
	assert.Equal(t, "/Users/michaelthornton", a.Next().String())
	assert.Equal(t, "/Users", a.Next().String())
	assert.Nil(t, a.Next())
	assert.Nil(t, a.Next())
}
