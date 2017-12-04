package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var uuids = []struct {
	id            string
	possiblyValid bool
	valid         bool
}{
	{"asdads", false, false},
	{"0X000000A0000Z0000Y00Y0X000000000000", false, false},
	{"00000000-0000-0000-0000-000000000000", true, false},
	{"0X000000-0000-0000-00Y0-000000000000", true, false},
	{"FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF", true, false},
	{"ffffffff-ffff-ffff-ffff-ffffffffffff", true, false},
	{"ED83F96D-D14E-4A81-8C7A-021DA055DEF5", true, true},
	{"a4128ee3-5fd5-4359-bae4-6c03354f9a02", true, true},
	{"28C1E1D8-7E76-4C57-A25A-4B8C81A3EB2A", true, true},
}

func TestUUID(t *testing.T) {
	for _, uuid := range uuids {
		possiblyValid := IsValidFast(uuid.id)
		if uuid.possiblyValid != possiblyValid {
			t.Errorf("%s should be recognized as possibltyValid=%t but is possibltyValid=%t", uuid.id, uuid.possiblyValid, possiblyValid)
		}
		valid := IsValid(uuid.id)
		if uuid.valid != valid {
			t.Errorf("%s should be recognized as valid=%t but is valid=%t", uuid.id, uuid.valid, valid)
		}
	}
}

func TestUUIDOsMatching(t *testing.T) {
	assert.True(t, MatchesOS("", "ios", "bla"), "an empty os should always match")
	assert.False(t, MatchesOS("android", "ios", "bla"), "different OSs shouldn't match")
	assert.False(t, MatchesOS("android", "android", "ED83F96D-D14E-4A81-8C7A-021DA055DEF5"), "correct OS but wrong user id shouldn't match")
	assert.True(t, MatchesOS("ios", "ios", "ED83F96D-D14E-4A81-8C7A-021DA055DEF5"), "correct OS but correct user id should match")
}
