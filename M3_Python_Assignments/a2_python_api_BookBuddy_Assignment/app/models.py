
from flask_sqlalchemy import SQLAlchemy

db = SQLAlchemy()

class Book(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String(100), nullable=False)
    author = db.Column(db.String(100), nullable=False)
    published_year = db.Column(db.Integer, nullable=False)
    genre = db.Column(db.String(50), nullable=False)

    def __repr__(self):
        return f"<Book {self.title}>"

    def to_dict(self):
        # Ensure fields are returned in a specific order
        return {
            "id": self.id,
            "title": self.title,
            "author": self.author,
            "published_year": self.published_year,
            "genre": self.genre
        }
