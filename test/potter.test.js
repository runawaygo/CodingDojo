require("mocha")
var expect = require("chai").expect
var _ = require("lodash")

var potter = require("../src/potter")

describe("Potter", function(){
  it("Should give 8 if buy one book", function(){
    var result = potter.cal([1,0,0,0,0])
    expect(result).eql(8)
  })

  it("Should give 16 if buy two book", function(){
    var result = potter.cal([2,0,0,0,0])
    expect(result).eql(16)
  })

  it("Should give 8*n if buy n book", function(){
    var result = potter.cal([5,0,0,0,0])
    expect(result).eql(8 * 5)
  })

  it("Should give 15.2 if buy one bookA and one bookB", function(){
    var result = potter.cal([1,1,0,0,0])
    expect(result).eql(15.2)
  })

  it("Should give 23.2 if buy 2 bookA and 1 bookB", function(){
    var result = potter.cal([2,1,0,0,0])
    expect(result).eql(23.2)
  })

  it("Should give 100.4 if 54321", function(){
    var result = potter.cal([5,4,3,2,1])
    expect(result).eql(100)

    var result = potter.cal([10,8,6,4,2])
    expect(result).eql(200)
  })
})
