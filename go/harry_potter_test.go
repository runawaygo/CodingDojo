package harry_potter

import (
	"math/rand"
	"testing"
)

func TestAnswer(t *testing.T) {
	t.Run("1 kind of books", func(t *testing.T) {
		t.Run("Should pay 8 Dollar when buy 1 Book", func(t *testing.T) {
			actual := Answer(1)
			assert(8.0, actual, t)
		})

		t.Run("Should pay 16 Dollar when buy 2 Book", func(t *testing.T) {
			actual := Answer(2)
			assert(16.0, actual, t)
		})

		t.Run("Should pay 8*n Dollar when buy n Book", func(t *testing.T) {
			n := float64(rand.Int31() >> 3)

			actual := Answer(n)
			expect := float64(8 * n)

			assert(actual, expect, t)
		})
	})

	t.Run("2 kinds of books", func(t *testing.T) {
		t.Run("Should pay 15.2 Dollar when buy [1, 1] Book", func(t *testing.T) {

			actual := Answer(1, 1)
			expect := float64((8 + 8) * 0.95)
			assert(actual, expect, t)
		})

		t.Run("Should pay 23.2 Dollar when buy [1,2] Book", func(t *testing.T) {

			actual := Answer(1, 2)
			expect := float64((8+8)*0.95 + 8)
			assert(actual, expect, t)
		})

	})

	t.Run("3 kinds of books", func(t *testing.T) {
		t.Run("Should pay 21.6 Dollar when buy [1, 1, 1] Book", func(t *testing.T) {

			actual := Answer(1, 1, 1)
			expect := 8 * 3 * 0.9
			assert(actual, expect, t)
		})

		t.Run("Should pay 23.2 + 21.6 Dollar when buy [1,2,3] Book", func(t *testing.T) {

			actual := Answer(1, 2, 3)
			expect := float64(23.2 + 21.6)
			assert(actual, expect, t)
		})

	})

	t.Run("4 or 5 kinds of books", func(t *testing.T) {
		t.Run("Should pay 25.6 Dollar when buy [1, 1, 1, 1] Book", func(t *testing.T) {

			actual := Answer(1, 1, 1, 1)
			expect := 8 * 4 * 0.8
			assert(actual, expect, t)

		})

		t.Run("Tests set", func(t *testing.T) {

			assert(Answer(1, 1, 1, 1, 1), 8*5*0.75, t)
			assert(Answer(1, 2, 3, 4, 5), 100.4, t)
		})

		// t.Run("Should pay 23.2 + 21.6 Dollar when buy [1,2,3] Book", func(t *testing.T) {

		// 	actual := Answer(1, 2, 3)
		// 	expect := float64(23.2 + 21.6)
		// 	assert(actual, expect, t)
		// })

	})

	t.Run("Best discount", func(t *testing.T) {

	})
}

func assert(actual float64, expect float64, t *testing.T) {
	if actual != expect {
		t.Errorf("Expected:%f, Actual:%f", expect, actual)
	}
}
