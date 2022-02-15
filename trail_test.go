package trail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTrailPath = "/Users/name/personal/trail"
var testTrail = Trail{testTrailPath}
var testTrailPathNpr = "personal/trail"
var testTrailNpr = Trail{testTrailPathNpr}
var badTrail = Trail{""}

func TestComponents(t *testing.T) {
	_, err := testTrail.Components()
	assert.NoError(t, err)
}

func TestHasPhysicalRoot(t *testing.T) {
	tt := []struct {
		trail Trail
		hpr   bool
		err   bool
	}{
		{trail: testTrail, hpr: true, err: false},
		{trail: testTrailNpr, hpr: false, err: false},
		{trail: badTrail, hpr: false, err: true},
	}

	for _, test := range tt {
		hbr, err := hasPhysicalRoot(test.trail.asByteSlice())
		assert.Equal(t, test.hpr, hbr)
		assert.Equal(t, test.err, err != nil)
	}
}
