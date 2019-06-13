package bowling_fp

type Hehe int

func (a Hehe) double() Hehe {
	return a * 2
}

func Add(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
