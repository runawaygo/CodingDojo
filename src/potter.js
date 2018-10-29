var _ = require("lodash")

exports.cal = function(n){
  return 8 * 0.75 * 5 * n[4]
       + 8 * 0.8  * 4 * (n[3]-n[4])
       + 8 * 0.9  * 3 * (n[2]-n[3])
       + 8 * 0.95 * 2 * (n[1]-n[2])
       + 8 * 1    * 1 * (n[0]-n[1])
       - 0.4 * Math.min(n[4], n[2]-n[4])
}