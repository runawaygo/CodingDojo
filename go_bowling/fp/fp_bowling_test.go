package fp_bowling

import (
	"testing"
)

func TestThrow(t *testing.T) {
	t.Run("Throw", func(t *testing.T) {
		game := BuildNewGame()
		game.Throw(10)

		assert(10, game.score, t)

		game.Throw(10)
		assert(30, game.score, t)
	})

	t.Run("Throw 10", func(t *testing.T) {

		throws := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
		game := BuildNewGame()
		game.batchThrow(throws)
		assert(300, game.score, t)
	})

	t.Run("Throw with second throw", func(t *testing.T) {
		throws := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		game := BuildNewGame()
		game.batchThrow(throws)

		assert(20, game.score, t)

	})

	t.Run("Throw with second throw 2", func(t *testing.T) {
		throws := []int{5,3,5,3,5,3,5,3,5,3,5,3,5,3,5,3,5,3,5,3}
		game := BuildNewGame()
		game.batchThrow(throws)

		assert(80, game.score, t)
	})

	t.Run("Throw 2 throw to spare", func(t *testing.T) {
		game := BuildNewGame()
		game.batchThrow([]int{1, 9, 1, 2})

		assert(14, game.score, t)
	})


	t.Run("Throw spare end", func(t *testing.T) {

		throws := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 5, 10}
		game := BuildNewGame()
		game.batchThrow(throws)
		assert(20, game.score, t)
	})

	t.Run("Throw strike end", func(t *testing.T) {

		throws := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10, 10}
		game := BuildNewGame()
		game.batchThrow(throws)
		assert(30, game.score, t)
	})


	t.Run("Throw strike", func(t *testing.T) {
		game := BuildNewGame()
		game.batchThrow([]int{10, 1, 1})

		assert(14, game.score, t)
	})
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

func (game *Game) batchThrow(throws []int) {
	for i := 0; i < len(throws); i++ {
		game.Throw(throws[i])
	}
}
