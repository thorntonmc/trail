package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAncestor(t *testing.T) {
	p := Path("/Users/michaelthornton/personal/ancestor")
	a := p.Ancestors()

	assert.Equal(t, "/Users/michaelthornton/personal/ancestor", string(a.Next()))
	assert.Equal(t, "/Users/michaelthornton/personal", string(a.Next()))
	assert.Equal(t, "/Users/michaelthornton", string(a.Next()))
	assert.Equal(t, "/Users", string(a.Next()))
	assert.Equal(t, "/", string(a.Next()))
	assert.Equal(t, "/", string(a.Next()))
	assert.Equal(t, "/", string(a.Next()))
	assert.Equal(t, "/", string(a.Next()))
	assert.Equal(t, "/", string(a.Next()))
	assert.Equal(t, "/", string(a.Next()))
	assert.Equal(t, "/", string(a.Next()))

}
