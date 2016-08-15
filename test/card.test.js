require("mocha")
var expect = require("chai").expect
var _ = require("lodash")


describe("Card", function(){
  var Card = require("../src/card")
  it("HighCard", function(){
    var card = new Card([11,2,3,4,1])
    expect(card.level()).to.be.equal(0)
  })
  it("Pair",function(){
    var card = new Card([11,11,3,4,1])
    expect(card.level()).to.be.equal(1)
  })

  it("2Pair",function(){
    var card = new Card([11,11,4,4,2])
    
    expect(card.level()).to.be.equal(2)
    expect(card.cards).to.be.eql([1,4,4,11,11])
  })

  it("ThreeKind",function(){
    var card = new Card([11,11,11,4,1])
    
    expect(card.level()).to.be.equal(3)
  })

  it("Straight",function(){
    var card = new Card([11,10,9,8,7])
    
    expect(card.level()).to.be.equal(4)
  })  
})
