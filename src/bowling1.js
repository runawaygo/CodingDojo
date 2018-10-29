var _ = require("lodash")

class Turn {
  constructor(){
    this.action = this.firstThrow
    this.pass = false
    this.score = 0
  }
  throw(t){
    if(this.action) {
      this.action = this.action(t)      
    }
  }
  firstThrow(t){
    this.score = t
    if(t === 10){
      return this.afterStrike
    }
    return this.secondThrow(t)
  }
  secondThrow(t1){
    return (t2) => {
      this.score += t2
      if((t1+t2) === 10) {
        return this.afterSpare
      }
      this.pass = true
    }
  }
  afterStrike(t){
    this.score += t
    return this.afterStrike2
  }
  afterStrike2(t){
    this.score += t
    this.pass = true
  }
  afterSpare(t){
    this.score += t
    this.pass = true
  }
}

class Bowling {
  constructor(){
    this.currentTrun = 0
    this.turns = []
    this.subTurns = []
    this.action = this.firstThrow
  }
  bThrow(throws){
    for(let i = 0;i<throws.length;i++){
      this.throw(throws[i])
    }
  }

  throw(t){
    this.action = this.action(t)
  }
  get score(){
    let count = 0
    for(let i = 0;i<10;i++){
      count += this.turns[i].score
    }
    return count
  }

  newTurn(){
    this.currentTrun++
    const turn = new Turn()
    this.turns.push(turn)
    this.subTurns.unshift(turn)
  }
  pubToTurn(t){
    for(let i = this.subTurns.length-1;i>=0;i--){
      const turn = this.subTurns[i]
      turn.throw(t)

      if(turn.pass){
        this.subTurns.pop()
      }
    }
  }
  firstThrow(t){
    this.newTurn()
    this.pubToTurn(t)

    if(t === 10) {
      if(this.currentTrun === 10) return this.endThrow

      return this.firstThrow
    }

    return this.secondThrow

  }

  secondThrow(t){
    this.pubToTurn(t)
    if(this.currentTrun === 10) 
      return this.endThrow

    return this.firstThrow
  }

  endThrow(t){
    this.pubToTurn(t)
    return this.endThrow
  }
}

module.exports = Bowling


