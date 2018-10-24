var _ = require("lodash")

exports.cal = function(n){
  if(arguments.length === 2){
    var bookACount = arguments[0]
    var bookBCount = arguments[1]

    return 8 * 0.95 * 2 * bookBCount + 8 * (bookACount - bookBCount)
  }

  return  n?8*n:0
}