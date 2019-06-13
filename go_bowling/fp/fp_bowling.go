package fp_bowling

type Game struct {
	score        int
	turns        int
	isEnd        bool
	throwHandler func(int) interface{}
}

func BuildNewGame() *Game{
	game := new(Game)
	game.score = 0
	game.turns = 0

	game.throwHandler = game.firstThrow
	return game
}
func (game *Game) Throw(n int) *Game{
	println(game.turns)
	println(game.score)
	println(n)
	println("-----------")
	game.throwHandler = game.throwHandler(n).(func(int)interface{})
	return game
}

func (game *Game) firstThrow(n int) interface{}{
	if game.turns==10 {
		return nil
	}

	game.turns++

	game.score+=n

	if n==10 {
		if game.turns == 10 {
			return game.afterStrikeEndAppend
		}
		return game.afterStrike
	}

	return game.secondThrow(n)
}

func (game *Game) secondThrow(n1 int) (func(int)interface{}){
	return func(n2 int) interface{}{
		game.score += n2

		if n1+n2==10 {
			if game.turns == 10 {
				return game.afterSpareEndAppend
			}
			return game.afterSpare
		}

		return game.firstThrow
	}
}

func(game *Game) afterSpare(n int) interface{}{
	game.score += n
	return game.firstThrow(n)
}

func(game *Game) afterStrike(n int) interface{}{
	if game.turns==10 {
		return nil
	}

	game.turns++

	game.score += n
	game.score += n

	if n==10 {
		if game.turns == 10 {
			return game.afterSpareEndAppend
		}
		return game.afterStrikeStrike
	}

	return game.afterStrike2(n)

}

func(game *Game) afterStrikeStrike(n int) interface{}{
	if game.turns==10 {
		return nil
	}

	game.turns++

	game.score += n
	game.score += n
	game.score += n

	if n==10 {
		if game.turns == 10 {
			return game.afterStrikeStrikeEndAppend
		}
		return game.afterStrikeStrike
	}

	return game.afterStrike2(n)
}

func(game *Game) afterStrike2(n1 int) (func(int)interface{}){
	return func(n2 int)interface{}{
		game.score += n2

		game.score += n2
		if n1+n2 == 10 {
			return game.afterSpare
		}

		return game.firstThrow

	}
}

func(game *Game) afterStrikeEndAppend(n int) interface{}{
	game.score += n

	return func(n2 int) interface{}{
		game.score += n2
		return (func(int)interface{})(nil)
	}
}

func(game *Game) afterStrikeStrikeEndAppend(n int) interface{}{
	game.score += n
	game.score += n

	return func(n2 int) interface{}{
		game.score += n2
		return (func(int)interface{})(nil)
	}
}

func (game *Game) afterSpareEndAppend(n int) interface{}{
	game.score += n
	return (func(int)interface{})(nil)
}

