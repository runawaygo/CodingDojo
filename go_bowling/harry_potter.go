package bowling

import "fmt"

type Game struct {
	score  int
	isEnd  bool
	throws []int
}

func (game *Game) Throw(n int) interface{} {
	fmt.Print(game)
	return Throw
}

type ThrowHandler func(int) (interface{}, int, bool)

func Match() (interface{}, int, bool) {
	turns := 0
	score := 0
	isEnd := false

	firstThrow := (ThrowHandler)(nil)
	secondThrow := (ThrowHandler)(nil)
	afterSpare := (ThrowHandler)(nil)
	// afterStrike := (ThrowHandler)(nil)
	// afterStrike2 := (ThrowHandler)(nil)
	// afterStrikeStrike := (ThrowHandler)(nil)

	turnUp := func() {
		turns++
		if turns == 10 {
			isEnd = true
		}
	}

	firstThrow = func(t int) (interface{}, int, bool) {
		score += t
		if t != 10 {
			return secondThrow(t)
		} else {
			turnUp()

			returnHandler := afterSpare
			if isEnd {
				returnHandler = (ThrowHandler)(nil)
			}

			return returnHandler, score, isEnd
		}
	}

	secondThrow = func(t1 int) (interface{}, int, bool) {
		return ThrowHandler(func(t2 int) (interface{}, int, bool) {
			score += t2
			turnUp()

			returnHandler := (ThrowHandler)(nil)
			if isEnd != true {
				if t1+t2 != 10 {
					returnHandler = firstThrow
				} else {
					returnHandler = afterSpare
				}
			}

			return returnHandler, score, isEnd
		}), score, isEnd
	}

	afterSpare = func(t int) (interface{}, int, bool) {
		score += t
		return firstThrow(t)
	}

	// afterStrike = func(t int) (interface{}, int, bool) {
	// 	score += t
	// 	return firstThrow(t)
	// }

	// afterStrikeStrike = func(t int) (interface{}, int, bool) {
	// 	score += t
	// 	return firstThrow(t)
	// }

	// afterStrike2 = func(t int) (interface{}, int, bool) {
	// 	score += t
	// 	return firstThrow(t)
	// }

	return firstThrow, score, isEnd
}

func Throw(n int) interface{} {
	return Throw
}
