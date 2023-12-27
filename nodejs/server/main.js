const http = require('node:http');

const host = 'localhost';
const port = '8000'

const requestListener = function (req, res) {
    res.writeHead(200);
    res.end("const os = require('node:os');console.log(os.userInfo().username);");
};

const server = http.createServer(requestListener);
server.listen(port, host, () => {
    console.log(`Server is running on http://${host}:${port}`);
});