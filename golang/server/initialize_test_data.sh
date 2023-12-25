curl -v -F "pluginname=whoami" -F file=@../plugin_whoami/whoami.so http://localhost:3333/addplugin/
curl -v -X POST -d "pluginname=whoami" -d "implantname=`hostname`" http://localhost:3333/linkimplantplugin/
