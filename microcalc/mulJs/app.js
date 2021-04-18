const express = require('express');
const app = express();
const morgan = require('morgan');
const debug = require('debug')('app');
const router = express.Router()
const bodyParser = require('body-parser');
const APP_VERSION = 'v1';
const APP_BASEPATH = `/api/${APP_VERSION}`;
const APP_SERVICE = `name: mul, version: ${APP_VERSION}`;
const APP_PORT = process.env.APP_PORT || 8083;

app.get('/', (req, res) => {
  res.send(`welcome all - ${APP_SERVICE}`);
});

router
    .get('/status', (_, res) => res.status(200).json({ data: 'ok' }))
    .post('/mul', (req, res) => {
        console.log("req.body = ",req.body)
        if (!req.body
            || !req.body.operands
            || !Array.isArray(req.body.operands)
            || req.body.operands.length !== 2) {
            return res.status(400).json({ error: 'Invalid input', service: APP_SERVICE });
        }

        let [op1, op2] = req.body.operands;
        debug('Received', op1, op2);

    });

app.use(morgan('combined'))
app.use(bodyParser.json());
app.use(APP_BASEPATH, router);
app.use((_, res) => {
    res.status(404).end();
});

app.listen(APP_PORT, () => {
  console.log(`service ${APP_SERVICE} launched at  ${APP_PORT}`);
});
