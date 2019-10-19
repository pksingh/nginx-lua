# Commandline Calculator 
A minimalstic calculator.

## Features
- Only supported basic operations : [ +, -, *, / ]
- Also supports long chanin of expressions: like `1 + 2 * 3 - 4 / 2` on the `calc>` prompt.

## Building
 We need to fire `go build cmdcalc.go` from the console inside the parent folder.
 This will generate `cmdcalc.exe` or `cmdcalc` binary files based on the platform you are using.

## Running
This will be pretty straight forward. We need to call the binary file generated in the build stage.
In our case; `cmdcalc.exe` and we can play with few inputs and check.
```
C:\Users\Home\Desktop\uservices>go build cmdcalc.go

C:\Users\Home\Desktop\uservices>cmdcalc.exe
calc>1 + 2
3
calc>1 + 2 + 3
6
calc>1 + 2 + 3 - 1
5
calc>1 + 2 + 3 - 1 * 2
10
calc>1 + 2 + 3 - 1 * 2 / 5
2
calc>1 + 2 + 3 ^ 1 * 2 / 5  
error: operator not from ( + - * / )
calc>calc>^C

```
