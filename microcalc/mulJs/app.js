const express = require('express');
const app = express();
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

app.use(APP_BASEPATH, router);

app.listen(APP_PORT, () => {
  console.log(`service ${APP_SERVICE} launched at  ${APP_PORT}`);
});
