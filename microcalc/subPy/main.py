from flask import Flask, request
from flask_restful import Resource, Api, abort

APP_VERSION = 'v1'
APP_BASEPATH = '/api/' + APP_VERSION
APP_SERVICE = 'name: sub, version: ' + APP_VERSION
APP_PORT = '8082'

app = Flask(__name__)
@app.route('/')
def hello():
    return {'world': 'welcome all : '+APP_SERVICE}

api = Api(app)

class GetStatus(Resource):
    def get(self):
        return {'data': 'ok'}

class PostSub(Resource):
    def post(self):
        body = request.get_json()
        ops = body['operands'] if body is not None else []
        print('Received body: {}'.format(body))

        if (len(ops) != 2 or int(ops[1]) == 0):
            return abort(400, error='Invalid Input', service=APP_SERVICE)

        op1,op2 = ops
        print('Received {} {}'.format(op1, op2))

        op1 = int(op1)
        op2 = int(op2)

        return {
            'result': int(op1-op2),
            'operands': [op1,op2],
            'service': APP_SERVICE
        }

api.add_resource(GetStatus, APP_BASEPATH + '/status')
api.add_resource(PostSub, APP_BASEPATH + '/sub')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=APP_PORT, debug=True)
