const express = require('express');
const app = express();
const APP_VERSION = 'v1';
const APP_BASEPATH = `/api/${APP_VERSION}`;
const APP_SERVICE = `name: mul, version: ${APP_VERSION}`;
const APP_PORT = process.env.APP_PORT || 8083;

app.get('/', (req, res) => {
  res.send('Hello World!');
});

app.listen(APP_PORT, () => {
  console.log(`service ${APP_SERVICE} launched at  ${APP_PORT}`);
});
