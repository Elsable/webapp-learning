const s = 'Keep On Asking, and It Will Be Given You'
const r = s.replace(/[a-z]+/g, function (m) {
  return m.toUpperCase()
})

console.log(r)

const arr = [100, 1, 20, 43, 30, 11, 4]
arr.sort((a, b) => { return b - a })
console.log(arr)
