# Modular C2

A C2 framework using plugin functionality in golang. 

The goal is to have a c2 that is very easy to extend on the fly.

# Golang version

** NOTE: This does not work on windows since the plugin is not implemented on windows **

## Usage
Start the server:
```
go run main.go <listen host>
```

Run the implant:
```
go run main.go <listenhost>+":3333" "hemmeligt_kodeord"
```

## Create plugins
Look at the whoami. The implant will all the `DoIt` method.

All plugins will execute in turn, so long running plugins would need to use go-routines.
