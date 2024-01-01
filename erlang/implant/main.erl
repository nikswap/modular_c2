-module(main).
-export([start/0]).

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

start() ->
    run_commands(['../plugins/test.erl']).

% loop with
% heartbeat
% download plugins to temp files and save file names
% run commands
