var _ = require("lodash")

class BowlingSM {
  constructor(){
    this.currentFrameIndex = 0
    this.frames = [0,0,0,0,0,0,0,0,0,0]
    this.action = this.firstThrow
  }

  get score(){
    let count = 0
    for(let i = 0;i<10;i++){
      count += this.frames[i]
    }
    return count
  }

  bThrow(throws){
    for(let i = 0;i<throws.length;i++){
      
      
      this.throw(throws[i])
    }
  }

  throw(t){
    
    this.action = this.action(t)
    return this
  }

  firstThrow(t){
    if(t === 10){
      this.frames[this.currentFrameIndex] = t
      this.currentFrameIndex++
      return this.afterStrike
    }
    return this.secondThrow(t)
  }
  secondThrow(t1){
    return (t2) => {
      this.frames[this.currentFrameIndex] = t1 + t2
      this.currentFrameIndex++

      if((t1+t2) === 10) {
        return this.afterSpare
      }

      return this.firstThrow
    }
  }

  afterSpare(t){
    this.frames[this.currentFrameIndex-1] += t

    if(t === 10){
      return this.afterStrike
    }

    return this.secondThrow(t)
  }
  afterStrike(t){
    this.frames[this.currentFrameIndex-1] += t    

    if(t === 10){
      this.frames[this.currentFrameIndex] = t
      this.currentFrameIndex++ 

      return this.afterStrikeStrike
    }
    return this.afterStrike2(t)

  }

  afterStrike2(t1){
    return (t2) => {
      this.frames[this.currentFrameIndex-1] += t2

      this.frames[this.currentFrameIndex] = t1 + t2
      this.currentFrameIndex++

      return (t1+t2) === 10 ? this.afterSpare : this.firstThrow
    }
  }
  afterStrikeStrike(t){
    this.frames[this.currentFrameIndex - 2] += t
    this.frames[this.currentFrameIndex - 1] += t

    if(t === 10){
      this.frames[this.currentFrameIndex] = t
      this.currentFrameIndex++
      return this.afterStrikeStrike
    }

    return this.afterStrike2(t)
  }
}

class Frame {
  constructor(){
    this.throws = []
    this.pass = false
  }
  get score(){
    return _.sum(this.throws)
  }
  get isFull(){
    if(this.throws[0] === 10 && this.throws.length >= 3)
      return true

    if((this.throws[0] + this.throws[1]) === 10 && this.throws.length >= 3)
      return true

    if((this.throws[0] + this.throws[1]) < 10 && this.throws.length >= 2) 
      return true

    return false
  }
  throw(t){
    if(this.isFull) return

    this.throws.push(t)

    if(t === 10){
      this.pass = true
    }
    if(this.throws.length === 2){
      this.pass = true
    }
  }
}

class Bowling {
  constructor(){
    this.currentFrameIndex = 0
    this.frames = [new Frame()]
  }

  get score(){
    let count = 0
    for(let i = 0;i<10;i++){
      count += this.frames[i].score
    }
    return count
  }
  get currentFrame(){
    return this.frames[this.currentFrameIndex]
  }

  bThrow(throws){
    for(let i = 0;i<throws.length;i++){
      
      
      this.throw(throws[i])
    }
  }

  throw(t){
    if(this.currentFrameIndex>=2) this.frames[this.currentFrameIndex-2].throw(t)

    if(this.currentFrameIndex>=1) this.frames[this.currentFrameIndex-1].throw(t) 

    this.currentFrame.throw(t)
    if(this.currentFrame.pass){
      this.currentFrameIndex++
      this.frames.push(new Frame())
    }

    
    
  }
}

// module.exports = BowlingSM
module.exports = Bowling


