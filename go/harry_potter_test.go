package harry_potter

import (
	"math/rand"
	"testing"
)

func TestAnswer(t *testing.T) {
	t.Run("Should pay 8 Dollar when buy 1 Book", func(t *testing.T) {
		result := answer(1)
		if result != 8 {
			t.Errorf("Expected:%d, Actual:%f", 8, result)
		}
	})

	t.Run("Should pay 16 Dollar when buy 2 Book", func(t *testing.T) {
		result := answer(2)
		if result != 16 {
			t.Errorf("Expected:%d, Actual:%f", 16, result)
		}

	})

	t.Run("Should pay 8*n Dollar when buy n Book", func(t *testing.T) {
		n := rand.Int31() >> 3

		result := answer(n)
		expect := float64(8 * n)
		if result != expect {
			t.Errorf("Expected:%f, Actual:%f", expect, result)
		}
	})
}
