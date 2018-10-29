require("mocha")
var expect = require("chai").expect
var _ = require("lodash")

describe("Bowling SM", function(){
  var Bowling = require("../src/bowling1")
  it("Normal1", function(){
    const bowling = new Bowling()
    bowling.bThrow([2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2])
    expect(bowling.score).to.eql(40)
  })

  it("Normal2", function(){
    const bowling = new Bowling()
    bowling.bThrow([2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,0,2])
    expect(bowling.score).to.eql(36)
  })

  it("Spare1", function(){
    const bowling = new Bowling()
    bowling.bThrow([2,8,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2])
    expect(bowling.score).to.eql(48)
  })

  it("Spare2", function(){
    const bowling = new Bowling()
    bowling.bThrow([2,8,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,8,10])
    expect(bowling.score).to.eql(64)
  })

  it("Strike1", function(){
    const bowling = new Bowling()
    bowling.bThrow([10,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2])
    expect(bowling.score).to.eql(50)
  })

  it("Strike2", function(){
    const bowling = new Bowling()
    bowling.bThrow([10,10,10,10,10,10,10,10,10,10,10,10])
    expect(bowling.score).to.eql(300)
  })
})
