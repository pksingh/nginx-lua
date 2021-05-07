const Koa = require("koa");             // import koa
const logger = require("koa-logger");   // import koa logger

// define the APP constants
const APP_VERSION = 'v1';
const APP_BASEPATH = `/api/${APP_VERSION}`;
const APP_SERVICE = `name: div, version: ${APP_VERSION}`;
const APP_PORT = process.env.APP_PORT || 8084;

//create the koa application
const app = new Koa();
app.use(logger());                      //Enable logger

app.use(async ctx => {
  ctx.body = 'Hello World';
});

// Turn on the server
app.listen(APP_PORT, () => {
  console.log(`service ${APP_SERVICE} launched at  ${APP_PORT}`);
});
