from flask import Flask,request
from users import users, companies


app = Flask(__name__)

@app.route('/')
def hello():
    suma = 24+222
    return {'suma':suma}

@app.route('/users')
def handleUsers():
    return {"users":companies}

@app.route('/multiply',methods=['POST'])
def multiply():
    req_data = request.get_json()
    num1 = req_data['num1']
    num2 = req_data['num2']
    
    return {"result":num1*num2}

if __name__ == '__main__':
    app.run(host="0.0.0.0",port=4000, debug=True)