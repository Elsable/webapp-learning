const fs = require('fs')

function readFile_pr (fname) {
  return new Promise((resolve) => {
    fs.readFile(fname, 'utf-8', (err, s) => {
      resolve(s)
    })
  })
}

readFile_pr('../../text/a.txt')
  .then((text) => {
    console.log('a.txtを読みこみました', text)
    return readFile_pr('../../text/b.txt')
  })
  .then((text) => {
    console.log('b.txtを読みこみました', text)
    return readFile_pr('../../text/c.txt')
  })
  .then((text) => {
    console.log('c.txtを読みこみました', text)
  })
