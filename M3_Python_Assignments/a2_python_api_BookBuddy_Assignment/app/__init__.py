from flask import Flask

# from flask_cors import CORS

from .config import Config
from .models import db
from .book_routes import main_routes

def create_app():
    app = Flask(__name__)
    app.config.from_object(Config)  # Ensure this path is correct
    db.init_app(app)    # Initializes the database with Flask app

    # CORS(app)
    
    app.register_blueprint(main_routes) # Register routes (if applicable)
    return app