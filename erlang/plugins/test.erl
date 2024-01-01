-module(test).
-export([start/0,get_result/0]).
-on_load(start/0).

start() ->
    io:format("HI FROM TEST PLUGIN~n").

get_result() ->
    io:format("Running code~n"),
    os:cmd("whoami").
