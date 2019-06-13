package bowling_fp

type Match struct {
	score int
	isEnd bool
}

func NewMatch() *Match {
	match := new(Match)
	match.score = 0
	match.isEnd = false
	return match
}

func (match *Match) Throw(t int) {
	match.score += t
}
