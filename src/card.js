var _ = require("lodash")

function Card(cards){
  this.cards = _.sortBy(cards)
}

Card.prototype.level = function(){
  if(this.Straight()) return 4
  if(this.threeKind()) return 3
  if(this.twoPair()) return 2
  if(this.pair()) return 1

  return 0
}

Card.prototype.Straight = function(){
  for(var i=4;i>0;i--){
    if((this.cards[i]-this.cards[i-1]) != 1) return false
  }  
  return true
}

Card.prototype.threeKind = function(){
  for(var i=4;i>1;i--){
    if(this.cards[i]==this.cards[i-1] && this.cards[i]==this.cards[i-2]){
      var card = this.cards[i]
      this.cards.splice(i-2,3)
      this.cards.push(card, card, card)
      return true
    }
  }  
  return false
}

Card.prototype.twoPair = function(){
  var pairs = []
  var single;

  for(var i=4;i>=0;i--){
    if(i==0 || this.cards[i]!=this.cards[i-1] ) {
      single = this.cards[i]
    }
    else{
      pairs.push(this.cards[i])
      i--
    }
  }
  if(pairs.length == 2) {
    this.cards = [single, pairs[1], pairs[1], pairs[0], pairs[0]]
    return true
  }
  return false
}

Card.prototype.pair = function(){
  for(var i=4;i>0;i--){
    if(this.cards[i]==this.cards[i-1]){
      var card = this.cards[i]
      this.cards.splice(i-1,2)
      this.cards.push(card, card)
      return true
    }
  }  
  return false
}

module.exports = Card