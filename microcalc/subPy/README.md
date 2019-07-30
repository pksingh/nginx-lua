# SUB service /for Calculator/ 
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
 * Running on http://0.0.0.0:8082/ (Press CTRL+C to quit)
127.0.0.1 - - [26/Jul/2019 20:47:28] "GET / HTTP/1.1" 200 -
127.0.0.1 - - [26/Jul/2019 20:48:30] "GET /status HTTP/1.1" 404 -
127.0.0.1 - - [26/Jul/2019 20:50:20] "GET /api/v1/status HTTP/1.1" 200 -
Received body: None
127.0.0.1 - - [26/Jul/2019 20:51:29] "POST /api/v1/sub HTTP/1.1" 200 -

```

These are the corresponding curl commands we have fired and the response as follows.

```
C:\Users\Home\Desktop\uservices>curl http://localhost:8082/       
{
  "world": "welcome all : name: sub, version: v1"
}

C:\Users\Home\Desktop\uservices>curl http://localhost:8082/status
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<title>404 Not Found</title>
<h1>Not Found</h1>
<p>The requested URL was not found on the server. If you entered the URL manually please check your spelling and try again.</p>

C:\Users\Home\Desktop\uservices>curl http://localhost:8082/api/v1/status
{
    "data": "ok"
}

C:\Users\Home\Desktop\uservices>curl -X POST http://localhost:8082/api/v1/sub -d "{ \"operands\": [3,1] }"
null
```

---

```
 * Restarting with stat
 * Debugger is active!
 * Debugger PIN: 230-188-425
 * Running on http://0.0.0.0:8082/ (Press CTRL+C to quit)
Received body: {'operands': [3, 1]}
Received 3 1
127.0.0.1 - - [27/Jul/2019 21:37:57] "POST /api/v1/sub HTTP/1.1" 200 -
Received body: {'operands': '[3,1]'}
127.0.0.1 - - [27/Jul/2019 21:40:17] "POST /api/v1/sub HTTP/1.1" 400 -
```

```
C:\Users\Home\Desktop\uservices>curl -X POST http://localhost:8082/api/v1/sub -H "Content-Type: application/json" -d "{ \"operands\": [3,1] }"
null

C:\Users\Home\Desktop\uservices>curl -X POST http://localhost:8082/api/v1/sub -H "Content-Type: application/json" -d "{ \"operands\": \"[3,1]\" }"
{
    "error": "Invalid Input",
    "service": "name: sub, version: v1"
}

C:\Users\Home\Desktop\uservices>
```

---

```
 * Restarting with stat
 * Debugger is active!
 * Debugger PIN: 230-188-425
 * Running on http://0.0.0.0:8082/ (Press CTRL+C to quit)
127.0.0.1 - - [28/Jul/2019 09:45:30] "POST /api/v1/sub HTTP/1.1" 400 -
Received body: {'operands': [3, 1, 1]}
127.0.0.1 - - [28/Jul/2019 09:46:00] "POST /api/v1/sub HTTP/1.1" 400 -
Received body: {'operands': [3, 1]}
Received 3 1
127.0.0.1 - - [28/Jul/2019 09:46:07] "POST /api/v1/sub HTTP/1.1" 200 -

```

```
C:\Users\Home\Desktop\uservices>curl -X POST http://localhost:8082/api/v1/sub -H "Content-Type: application/json" -d "{ \"operands\": ['3',1] }" 
{
    "message": "Failed to decode JSON object: Expecting value: line 1 column 16 (char 15)"      
}

C:\Users\Home\Desktop\uservices>curl -X POST http://localhost:8082/api/v1/sub -H "Content-Type: application/json" -d "{ \"operands\": [3,1,1] }"
{
    "error": "Invalid Input",
    "service": "name: sub, version: v1"
}

C:\Users\Home\Desktop\uservices>curl -X POST http://localhost:8082/api/v1/sub -H "Content-Type: application/json" -d "{ \"operands\": [3,1] }"
{
    "result": 2,
    "operands": [
        3,
        1
    ],
    "service": "name: sub, version: v1"
}

C:\Users\Home\Desktop\uservices>
```
