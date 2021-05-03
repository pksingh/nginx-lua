const Koa = require("koa");             // import koa
const logger = require("koa-logger");   // import koa logger
const app = new Koa();

app.use(logger());                      //Enable logger
app.use(async ctx => {
  ctx.body = 'Hello World';
});

app.listen(4000, () => console.log("Server Listening on Port 4000"));
