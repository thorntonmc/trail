package trail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComponentRead(t *testing.T) {
	var c *Components
	// with root
	c, _ = testTrail.Components()
	assert.Equal(t, string(rootDir), string(c.Next()))
	assert.Equal(t, "Users", string(c.Next()))
	assert.Equal(t, "name", string(c.Next()))
	assert.Equal(t, "personal", string(c.Next()))
	assert.Equal(t, "trail", string(c.Next()))

	c, _ = testTrailNpr.Components()
	assert.Equal(t, "personal", string(c.Next()))
}
