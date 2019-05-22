package bowling

type Game struct {
	score        int
	turns        int
	isEnd        bool
	needAppend   bool
	throwHandler interface{}
}

func NewGame() *Game {
	game := new(Game)
	game.isEnd = false
	game.score = 0
	game.turns = 0
	game.needAppend = false
	game.throwHandler = game.firstThrow
	return game
}

func (game *Game) Throw(n int) *Game {
	if game.isEnd && !game.needAppend {
		panic("The game is end!!!")
	}
	game.throwHandler = game.throwHandler.(func(int) interface{})(n)
	return game
}

func (game *Game) upTurn() *Game {
	game.turns++
	if game.turns == 10 {
		game.isEnd = true
	}

	return game
}

func (game *Game) firstThrow(t int) interface{} {
	game.score += t

	if t != 10 {
		return game.secondThrow(t)
	}

	game.upTurn()

	if game.isEnd {
		game.needAppend = true
		return game.afterEndStrikeThrow
	}

	return game.afterStrike
}

func (game *Game) secondThrow(t1 int) interface{} {
	return func(t2 int) interface{} {
		game.score += t2

		game.upTurn()

		if t1+t2 != 10 {
			return game.firstThrow
		}

		if game.isEnd {
			game.needAppend = true
			return game.afterEndSpareThrow
		}

		return game.afterSpareThrow
	}
}

func (game *Game) afterSpareThrow(t int) interface{} {
	game.score += t
	return game.firstThrow(t)
}

func (game *Game) afterEndSpareThrow(t int) interface{} {
	game.score += t
	game.needAppend = false
	return nil
}

func (game *Game) afterStrike(t int) interface{} {
	game.score += t
	game.score += t

	if t != 10 {
		return game.afterStrike2(t)
	}

	game.upTurn()

	if game.isEnd {
		game.needAppend = true
		return game.afterEndStrikeThrow
	}

	return game.afterStrikeStrike
}

func (game *Game) afterStrike2(t1 int) interface{} {
	return func(t2 int) interface{} {
		game.score += t2
		game.score += t2

		game.upTurn()

		if t1+t2 != 10 {
			return game.firstThrow
		}

		return game.afterSpareThrow
	}
}

func (game *Game) afterStrikeStrike(t int) interface{} {
	game.score += t
	game.score += t
	game.score += t

	if t != 10 {
		return game.afterStrike2
	}

	game.upTurn()
	if game.isEnd {
		game.needAppend = true
		return game.afterEndStrikeStrikeThrow
	}

	return game.afterStrikeStrike
}

func (game *Game) afterEndStrikeThrow(t int) interface{} {
	game.score += t
	return func(t2 int) interface{} {
		game.score += t
		game.needAppend = false
		return nil
	}
}

func (game *Game) afterEndStrikeStrikeThrow(t int) interface{} {
	game.score += t
	return game.afterEndStrikeThrow(t)
}
