# MUL(tiplication) service /for Calculator/ 
This is simple math service; which can be used for multi purpose. (Mostly for demo)


## BUILD
Depedency libs are defined in the `package.json` file.  
We need to resolve them before app starting.  
Go to the same folder where 'package.json' is present and fire below command in the console:  
`node install`

## RUNNING

Make sure node is installed and all depedency have been available.  
We can verify that firing `node -v` or `npm -v` on the console.  

We can bring up the service by passing main/start js script file to node interpretor on the console.
`node app.js` or `npm start`


This will bring the service up and show the below.
```
C:\Users\Home\Desktop\uservices\microcalc\mulJs>npm start

> uservices.microcalc.mul@0.0.1 start
> node app.js

service launched at  3000
Terminate batch job (Y/N)? y
```

or  

In development; we can start application by running `npm run start:dev` which will auto restart if any changed found.

```
C:\Users\Home\Desktop\uservices\microcalc\mulJs>npm run start:dev

> uservices.microcalc.mul@0.0.1 start:dev
> pm2-dev app.js

===============================================================================
--- PM2 development mode ------------------------------------------------------
Apps started         : app
Processes started    : 1
Watch and Restart    : Enabled
Ignored folder       : node_modules
===============================================================================
app-0  | service launched at  3000
>>>>> [PM2 DEV] Stopping current development session
Terminate batch job (Y/N)? y
```

