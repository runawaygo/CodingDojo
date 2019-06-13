package bowling_fp

import (
	"testing"
)

func TestMatch(t *testing.T) {
	match := NewMatch()
	match.Throw(5)
	match.Throw(4)
	assert(match.score, 9)

	match.Throw(10)
	assert(match.score, 19)

	match.Throw(9)
	assert(match.score, 37)
}

func assert(expect int, actual int, t *testing.T) {
	if actual != expect {
		t.Errorf("Expected:%d, Actual:%d", expect, actual)
	}
}

func assertBool(expect bool, actual bool, t *testing.T) {
	if actual != expect {
		t.Errorf("Expected:%t, Actual:%t", expect, actual)
	}
}
