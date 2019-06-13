package bowling_fp

import "testing"

func TestCurring(t *testing.T) {

	x := Add(5)
	println(x(1))
	println(x(3))
}
