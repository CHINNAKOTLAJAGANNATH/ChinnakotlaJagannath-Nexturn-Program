
class Config:
    SQLALCHEMY_DATABASE_URI = 'sqlite:///books.db'  # SQLite Database URI, This should point to the correct path
    SQLALCHEMY_TRACK_MODIFICATIONS = False  # Disable Flask-SQLAlchemy event system