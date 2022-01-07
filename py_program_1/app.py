from flask import Flask

app = Flask(__name__)

hello_message = "Hello World 1 form python webapp"
@app.route('/')
def index():
    return hello_message

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8080)

