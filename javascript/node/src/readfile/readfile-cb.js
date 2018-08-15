const fs = require('fs')

fs.readFile('../../text/kakugen.txt', 'utf-8', function (err, data) {
  console.log(data)
})
