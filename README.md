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
# Calculator Web Service /MONOlithic/

```
C:\Users\Home\Desktop\uservices\monocalc>go build .

C:\Users\Home\Desktop\uservices\monocalc>monocalc.exe
2020/04/04 11:14:35 Server Starting on 8080
^C
```

```
C:\Users\Home\Desktop\uservices\monocalc>curl http://localhost:8080/status
{"data":"ok"}
C:\Users\Home\Desktop\uservices\monocalc>curl http://localhost:8080/calculate
{"data":"called /calculate"}
```

---

```
C:\Users\Home\Desktop\uservices\monocalc>curl -X POST http://localhost:8080/calculate
404 page not found

C:\Users\Home\Desktop\uservices\monocalc>curl http://localhost:8080/api/v1/calculate

C:\Users\Home\Desktop\uservices\monocalc>curl -X POST http://localhost:8080/api/v1/calculate
{"data":"called /calculate"}
C:\Users\Home\Desktop\uservices\monocalc>
```

---

C:\Users\Home\Desktop\uservices\monocalc>monocalc.exe
2020/04/12 10:04:01 Server Starting on 8080
Received on /calculate:
Received on /calculate: data
Received on /calculate: input-data
Received on /calculate:

```
C:\Users\Home\Desktop\uservices\monocalc>curl -X POST http://localhost:8080/api/v1/calculate -d "-"
{"data":"called /calculate "}
C:\Users\Home\Desktop\uservices\monocalc>curl -X POST http://localhost:8080/api/v1/calculate -d "{ \"input\": \"data\" }"
{"data":"called /calculate data"}
C:\Users\Home\Desktop\uservices\monocalc>curl -X POST http://localhost:8080/api/v1/calculate -d "{ \"input\": \"input-data\" }" 
{"data":"called /calculate input-data"}
C:\Users\Home\Desktop\uservices\monocalc>
C:\Users\Home\Desktop\uservices\monocalc>curl -X POST http://localhost:8080/api/v1/calculate -d "{ \"wrong\": \"input-data\" }" 
{"data":"called /calculate "}
C:\Users\Home\Desktop\uservices\monocalc>
```

---
```
C:\Users\Home\Desktop\uservices\monocalc>monocalc.exe
2020/04/19 11:11:53 Server Starting on 8080
Received on /calculate: input-data
Returened on /calculate: map[operands:[input-data] origins:input-data result:WIP service:name: monocalc, version: v1]
```

```
C:\Users\Home\Desktop\uservices\monocalc>
C:\Users\Home\Desktop\uservices\monocalc>curl -X POST http://localhost:8080/api/v1/calculate -d "{ \"wrong\": \"input-data\" }"
{"data":"called /calculate " }
C:\Users\Home\Desktop\uservices\monocalc>curl -X POST http://localhost:8080/api/v1/calculate -d "-"
{"error":"Invalid Input","service":"name: monocalc, version: v1"}
C:\Users\Home\Desktop\uservices\monocalc>
C:\Users\Home\Desktop\uservices\monocalc>curl -X POST http://localhost:8080/api/v1/calculate -d "{ \"input\": \"input-data\" }"
{"operands":"[input-data]","origins":"input-data","result":"WIP","service":"name: monocalc, version: v1"}
```
---
```
C:\Users\Home\Desktop\uservices\monocalc>monocalc.exe
2020/04/19 11:17:08 Server Starting on 8080
Received on /calculate: input-data
2020/04/19 11:17:13 char i not allowed

C:\Users\Home\Desktop\uservices\monocalc>monocalc.exe
2020/04/19 11:19:11 Server Starting on 8080
Received on /calculate: 1 2 3
tokens:  [{1 1} {1 2} {1 3}]
Returened on /calculate: map[operands:[1 2 3] origins:1 2 3 result:3 service:name: monocalc, version: v1]
^C
```
---
```
C:\Users\Home\Desktop\uservices\monocalc>
C:\Users\Home\Desktop\uservices\monocalc>curl -X POST http://localhost:8080/api/v1/calculate -d "{ \"input\": \"input-data\" }"
curl: (56) Recv failure: Connection was reset

C:\Users\Home\Desktop\uservices\monocalc>
C:\Users\Home\Desktop\uservices\monocalc>curl -X POST http://localhost:8080/api/v1/calculate -d "{ \"input\": \"1 2 3\" }"
{"operands":"[1 2 3]","origins":"1 2 3","result":3,"service":"name: monocalc, version: v1"}

C:\Users\Home\Desktop\uservices\monocalc>
```

---

```
C:\Users\Home\Desktop\uservices\monocalc>monocalc.exe
2020/05/16 18:12:47 Server Starting on 8080
Received on /calculate: 1+2-3
tokens:  [{1 1} {0 +} {1 2} {0 -} {1 3}]
Evaluate: 1 + 2 = 3
Evaluate: 3 - 3 = 0
Returened on /calculate: map[operands:[1+2-3] origins:1+2-3 result:0 service:name: monocalc, version: v1]    
Received on /calculate: 1 2 # 3
2020/05/16 18:14:38 char # not allowed

C:\Users\Home\Desktop\uservices\monocalc>
```

```
C:\Users\Home\Desktop\uservices>curl -X POST http://localhost:8080/api/v1/calculate -d "{ \"input\": \"1+2-3\" }"
{"operands":"[1+2-3]","origins":"1+2-3","result":0,"service":"name: monocalc, version: v1"}

C:\Users\Home\Desktop\uservices>curl -X POST http://localhost:8080/api/v1/calculate -d "{ \"input\": \"1 2 # 3\" }"
curl: (56) Recv failure: Connection was reset
C:\Users\Home\Desktop\uservices>
```