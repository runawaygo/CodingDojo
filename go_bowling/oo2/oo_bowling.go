package oo_bowling

type Frame struct {
	score 		int
	throws 		[]int
	state	    int // 0 结束，1 初始状态，2 接受了一个不是10的球，3 spare，4 strike， 5 strike2
	stateHandler 	func(int) interface{}
}

func BuildNewFrame() *Frame{
	frame := new(Frame)
	frame.score = 0
	frame.state = 1

	frame.throws = make([]int, 0 , 4)
	frame.stateHandler = frame.firstThrow

	return frame
}

func (this *Frame) Throw(n int) *Frame{
	this.throws = append(this.throws, n)
	this.score += n

	this.stateHandler = this.stateHandler(n).(func(int) interface{})
	return this
} 

func (this *Frame) firstThrow(n int) interface{}{

	if n == 10{
		this.state = 4
		return this.afterStrikeThrow
	}

	this.state = 2
	return this.secondThrow
}

func (this *Frame) secondThrow(n int) interface{}{
	if this.throws[0]+this.throws[1] == 10{

		this.state = 3
		return this.afterSpareThrow
	}

	this.state = 0
	return (func(int)interface{})(nil)
}

func (this *Frame) afterSpareThrow(n int) interface{}{
	this.state = 0
	return (func(int)interface{})(nil)
}

func (this *Frame) afterStrikeThrow(n int) interface{}{
	this.state = 5
	return this.afterStrikeThrow2
}

func (this *Frame) afterStrikeThrow2(n int) interface{}{
	this.state = 0
	return (func(int)interface{})(nil)
}

type Game struct {
	score        	int
	currentFrameIndex int
	isEnd        	bool
	frames 			[]*Frame
}

func BuildNewGame() *Game{
	game := new(Game)
	game.score = 0
	game.currentFrameIndex = 0
	game.frames = make([]*Frame, 0, 10)

	for i:=0;i<10;i++ {
		game.frames =append(game.frames, BuildNewFrame())
	}

	return game
}
func (this *Game) Throw(n int) *Game{
	for i := this.currentFrameIndex;i>=0;i--{
		if this.frames[i].state == 0 {
			continue
		}
		this.frames[i].Throw(n)
	}

	this.calScore()

	if this.isEnd {
		return this
	}

	currentFrame := this.frames[this.currentFrameIndex]
	if currentFrame.state ==0 || currentFrame.state>2{
		if this.currentFrameIndex == 9 {
			this.isEnd = true
			return this
		}
		this.currentFrameIndex++
	}

	
	return this
}
func (this *Game) calScore()*Game{
	this.score = 0
	for i:=0;i<=this.currentFrameIndex;i++ {
		this.score += this.frames[i].score
	}
	return this
}




