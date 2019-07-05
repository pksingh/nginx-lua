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


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=APP_PORT, debug=True)
