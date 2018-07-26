const fs = require('fs')

function read_gfn (g, fname) {
  fs.readFile(fname, 'utf-8', (err, data) => {
    g.next(data)
  })
}

const g = (function * () {
  const a = yield read_gfn(g, '../../text/a.txt')
  console.log(a)
  const b = yield read_gfn(g, '../../text/b.txt')
  console.log(b)
  const c = yield read_gfn(g, '../../text/c.txt')
  console.log(c)
})()

g.next()
