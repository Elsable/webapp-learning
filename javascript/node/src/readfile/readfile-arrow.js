const fs = require('fs')

fs.readFile('../../text/kakugen.txt', 'utf-8', (err, data) => {
  console.log(data)
})
