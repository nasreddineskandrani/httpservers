const express = require('express');
const multer = require('multer');
var cors = require('cors');
const https = require('https');
const fs = require('fs');
const port = 3000;

// https: //stackoverflow.com/questions/11744975/enabling-https-on-express-js
var key = fs.readFileSync('./selfsigned.key');
var cert = fs.readFileSync('./selfsigned.crt');
var options = {
    key: key,
    cert: cert
};

app = express()

app.use(cors());

const upload = multer({
    dest: 'uploads/'
});

app.post('/api/uploadfile', upload.single('file'), (req, res) => {
    console.log(`new upload = ${req.file.filename}\n`);
    console.log(req.file);
    res.json({
        msg: 'Upload Works'
    });
});

app.use(function (err, req, res, next) {
    console.log('This is the invalid field ->', err.field)
    next(err)
});

var server = https.createServer(options, app);

server.listen(port, () => {
    console.log("server starting on port : " + port)
});
