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
calc>1+2
error: missing arguments (pass space between number and symbols)

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

---
# Commandline Statical Calculator

A Statical calculator.

## Features

- It supports basic statistics : like [ Total, Min, Max, Avg ]
- Also supports extended statistics: like [ Mean, Median, Mode, StandardDeviation, Variance ]
- Also accepts user data from cmd loops: like `1, 2, 3.3, 4.42, 0.5, 6,      7.8,     9` on the `statc>` prompt.
    It will work with single/multi space(' ') and comma(',') as separator.
    Also supports mix of space(' ') and comma(',') as separator.

## Building

 We need to fire `go build cmdstatc.go` from the console inside the parent folder.
 This will generate `cmdstatc.exe` or `cmdstatc` binary files based on the platform you are using.

## Running

This will be pretty straight forward. We need to call the binary file generated in the build stage.
In our case; `cmdstatc.exe` and we can play with few inputs and check.
```
C:\Users\Home\Desktop\uservices>go build cmdstatc.go

C:\Users\Home\Desktop\uservices>cmdstatc.exe

:: Statistics ::
Data: [1 2.3 3 4.56 0.5 6 7.8 9]
Total: 34.160000
Mean: 4.270000
Median: 2.530000
Sorted Median: 3.780000
Max: 9.000000
Max Indices: [7]
Min: 0.500000
Min Indices: [4]
Variance: 8.538800
Standard Deviation: 2.922123
Modes: []
Mode Repeat Count: 0


C:\Users\Home\Desktop\uservices>cmdstatc.exe
statc>1, 2, 3.3, 4.42, 0.5, 6,      7.8,     9 
inp: 1, 2, 3.3, 4.42, 0.5, 6,      7.8,     9 
inp : 1,2,3.3,4.42,0.5,6,7.8,9
:: Statistics ::
 Data: [1 2 3.3 4.42 0.5 6 7.8 9]
Total: 34.020000
Mean: 4.252500
Median: 2.460000
Sorted Median: 3.860000
Max: 9.000000
Max Indices: [7]
Min: 0.500000
Min Indices: [4]
Variance: 8.605794
Standard Deviation: 2.933563
Modes: []
Mode Repeat Count: 0
statc>statc>^C

C:\Users\Home\Desktop\uservices>
```

---
