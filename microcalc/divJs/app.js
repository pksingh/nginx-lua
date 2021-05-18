const Koa = require("koa");             // import koa
const logger = require("koa-logger");   // import koa logger
const koaRouter = require("koa-router"); // import koa router
const koaBody = require('koa-body');    // import koa body parser

// define the APP constants
const APP_VERSION = 'v1';
const APP_BASEPATH = `/api/${APP_VERSION}`;
const APP_SERVICE = `name: div, version: ${APP_VERSION}`;
const APP_PORT = process.env.APP_PORT || 8084;

//create the koa application
const app = new Koa();
app.use(logger());                      //Enable logger


// create a router for building routes
const router = koaRouter();

// Index Route
// context is a combo of the node request/response objects
router
    .prefix(APP_BASEPATH)
    .get("/status", async (ctx) => {
        // The response is the value of the context body
        ctx.body = "{data: 'ok'}";
    })
    .post('/div', async (ctx) => {
        console.log("request.body = ",ctx.request.body);

        if (!ctx.request.body
            || !ctx.request.body.operands
            || !Array.isArray(ctx.request.body.operands)
            || ctx.request.body.operands.length !== 2) {
                ctx.throw(400, `{ error: 'Invalid input', service: ${APP_SERVICE} }`);
        }

    });

// Register routes
app.use(koaBody());
app.use(router.routes());

// Turn on the server
app.listen(APP_PORT, () => {
  console.log(`service ${APP_SERVICE} launched at  ${APP_PORT}`);
});
