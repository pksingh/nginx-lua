# ADD service /for Calculator/ 
This is simple math service; which can be used for multi purpose. (Mostly for demo)


## BUILD
Depedency libs are defined in the requirements.txt file.  
We need to resolve them before app starting.  
We can install them by folling command  
`pip install -r requirements.txt`

## RUNNING
Make sure Python is installed and should be > 2.7 version.  
We can verify that firing `py -V` or `python -V` on the console.  

We can bring up the service by passing main python script file to python interpretor on the console.
`python main.py` or `py main.py`  

or  

We can set the env variable then start application by running `flask run`
```
set FLASK_APP=main.py     ( $ export FLASK_APP=main.py )
set FLASK_ENV=development ( $ export FLASK_ENV=development )
```

This will bring the service up and show the below.
```
C:\Users\Home\Desktop\uservices\microcalc\addPy>py main.py
 * Serving Flask app "main" (lazy loading)
 * Environment: development
 * Debug mode: on
 * Restarting with stat
 * Debugger is active!
 * Debugger PIN: 230-188-425
 * Running on http://0.0.0.0:8081/ (Press CTRL+C to quit)
127.0.0.1 - - [07/Jul/2019 21:04:53] "GET / HTTP/1.1" 200 -
127.0.0.1 - - [07/Jul/2019 21:05:00] "GET /status HTTP/1.1" 404 -
127.0.0.1 - - [07/Jul/2019 21:05:14] "GET /api/v1/status HTTP/1.1" 200 -

```

These are the corresponding curl commands we have fired and the response as follows.

```
C:\Users\Home\Desktop\uservices>curl http://localhost:8081/       
{
  "world": "welcome all"
}

C:\Users\Home\Desktop\uservices>curl http://localhost:8081/status
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<title>404 Not Found</title>
<h1>Not Found</h1>
<p>The requested URL was not found on the server. If you entered the URL manually please check your spelling and try again.</p>

C:\Users\Home\Desktop\uservices>curl http://localhost:8081/api/v1/status
{
    "data": "ok"
}

```

---

```
 * Running on http://0.0.0.0:8081/ (Press CTRL+C to quit)
Received body: {'operands': [1, 3, 5]}
127.0.0.1 - - [11/Jul/2019 21:20:45] "POST /api/v1/add HTTP/1.1" 200 -
Received body: None
127.0.0.1 - - [11/Jul/2019 21:21:28] "POST /api/v1/add HTTP/1.1" 200 -

```

```
C:\Users\Home\Desktop\uservices>curl -X POST http://localhost:8081/api/v1/add -H "Content-Type: application/json" -d "{ \"operands\": [1,3] }"
null

C:\Users\Home\Desktop\uservices>curl -X POST http://localhost:8081/api/v1/add -d "{ \"operands\": [1,3] }"
null

C:\Users\Home\Desktop\uservices>
```

---

```
 * Restarting with stat
 * Debugger is active!
 * Debugger PIN: 230-188-425
 * Running on http://0.0.0.0:8081/ (Press CTRL+C to quit)
Received body: None
127.0.0.1 - - [13/Jul/2019 19:57:02] "POST /api/v1/add HTTP/1.1" 400 -

Received body: {'operands': [1, 3]}
Received 1 3
127.0.0.1 - - [14/Jul/2019 10:59:54] "POST /api/v1/add HTTP/1.1" 200 -
```

```
C:\Users\Home\Desktop\uservices>curl -X POST http://localhost:8081/api/v1/add -d "{ \"operands\": [1,3] }"
{
    "error": "Invalid Input",
    "service": "name: add, version: v1"
}

C:\Users\Home\Desktop\uservices>curl -X POST http://localhost:8081/api/v1/add -H "Content-Type: application/json" -d "{ \"operands\": [1,3] }"
null

C:\Users\Home\Desktop\uservices>
```

---
