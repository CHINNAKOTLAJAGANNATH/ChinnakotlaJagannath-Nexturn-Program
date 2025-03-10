# app.py
from flask import Flask
from routes.book_routes import book_blueprint
from routes.member_routes import member_blueprint
from routes.transaction_routes import transaction_blueprint

# Initialize Flask App
app = Flask(__name__)

# Register Blueprints
app.register_blueprint(book_blueprint, url_prefix='/books')
app.register_blueprint(member_blueprint, url_prefix='/members')
app.register_blueprint(transaction_blueprint, url_prefix='/transactions')

@app.route('/')
def home():
    return {"message": "Welcome to the Library Management System API!"}

if __name__ == '__main__':
    app.run(debug=True)
