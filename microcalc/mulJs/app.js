const express = require('express');
const app = express();
const morgan = require('morgan');
const debug = require('debug')('app');
const router = express.Router()
const APP_VERSION = 'v1';
const APP_BASEPATH = `/api/${APP_VERSION}`;
const APP_SERVICE = `name: mul, version: ${APP_VERSION}`;
const APP_PORT = process.env.APP_PORT || 8083;

app.get('/', (req, res) => {
  res.send(`welcome all - ${APP_SERVICE}`);
});

router
    .get('/status', (_, res) => res.status(200).json({ data: 'ok' }))

app.use(morgan('combined'))
app.use(APP_BASEPATH, router);
app.use((_, res) => {
    res.status(404).end();
});

app.listen(APP_PORT, () => {
  console.log(`service ${APP_SERVICE} launched at  ${APP_PORT}`);
});
