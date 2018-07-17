const fs = require('fs')

function readFileEx (fname) {
  return new Promise((resolve, reject) => {
    fs.readFile(fname, 'utf-8', (err, data) => {
      resolve(data)
    })
  })
}

async function readAll () {
  const a = await readFileEx('../../text/a.txt')
  console.log(a)
  const b = await readFileEx('../../text/b.txt')
  console.log(b)
  const c = await readFileEx('../../text/c.txt')
  console.log(c)
}

readAll()
