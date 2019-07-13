from flask import Flask, request
from flask_restful import Resource, Api, abort

APP_VERSION = 'v1'
APP_BASEPATH = '/api/' + APP_VERSION
APP_SERVICE = 'name: add, version: ' + APP_VERSION
APP_PORT = '8081'

app = Flask(__name__)
@app.route('/')
def hello():
    return {'world': 'welcome all'}

api = Api(app)

class GetStatus(Resource):
    def get(self):
        return {'data': 'ok'}

class PostAdd(Resource):
    def post(self):
        body = request.get_json()
        ops = body['operands'] if body is not None else []
        print('Received body: {}'.format(body))

        if (len(ops) != 2 or int(ops[1]) == 0):
            return abort(400, error='Invalid Input', service=APP_SERVICE)

        op1,op2 = ops
        print('Received {} {}'.format(op1, op2))

api.add_resource(GetStatus, APP_BASEPATH + '/status')
api.add_resource(PostAdd, APP_BASEPATH + '/add')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=APP_PORT, debug=True)
