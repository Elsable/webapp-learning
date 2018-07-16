const fs = require('fs');
const request = require('request');

request('http://uta.pw/shodou/img/28/214.png')
    .pipe(fs.createWriteStream('../figure/test.png'))
