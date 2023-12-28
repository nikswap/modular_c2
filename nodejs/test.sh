curl -v -X POST -d '{"implant_name":`hostname`,"plugin_name":"whoami"}' localhost:8000/linkimplantplugin/
curl -v -X POST -d '{"plugin_content":"console.log(1+2);","plugin_name":"add"}' localhost:8000/addplugin/
curl -v -X POST -d '{"implant_name":`hostname`,"plugin_name":"add"}' localhost:8000/linkimplantplugin/
