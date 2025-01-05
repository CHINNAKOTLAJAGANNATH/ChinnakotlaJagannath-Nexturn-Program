from flask import Blueprint, request, jsonify
from .models import db, Book

main_routes = Blueprint('main_routes', __name__)

VALID_GENRES = ["Fiction", "Non-Fiction", "Mystery", "Sci-Fi"]

# Home route
@main_routes.route('/', methods=['GET'])
def welcome():
    return jsonify({"message": "Welcome to BookBuddy: A Book Collection Manager!"}), 200

# Add a new book
@main_routes.route('/books', methods=['POST'])
def add_book():
    data = request.get_json()   # Get JSON data from the request body
    title = data.get('title')
    author = data.get('author')
    published_year = data.get('published_year')
    genre = data.get('genre')

    if not title or not author or not published_year or not genre:
        return jsonify({"error": "Invalid data", "message": "All fields are required"}), 400
    
    if genre not in VALID_GENRES:
        return jsonify({"error": "Invalid genre", "message": f"Genre must be one of {VALID_GENRES}"}), 400
    
    if not isinstance(published_year, int) or published_year < 1000 or published_year > 9999:
        return jsonify({"error": "Invalid year", "message": "Published year must be a valid 4-digit year"}), 400

    new_book = Book(
        title=title, 
        author=author, 
        published_year=published_year, 
        genre=genre
    )

    db.session.add(new_book)
    db.session.commit()

    return jsonify({"message": "Book added successfully", "book_id": new_book.id}), 201

# Retrieve all books or filter by genre/author
@main_routes.route('/books', methods=['GET'])
def get_books():
    genre = request.args.get('genre')
    author = request.args.get('author')

    query = Book.query
    if genre:
        query = query.filter_by(genre=genre)
    if author:
        query = query.filter(Book.author.like(f"%{author}%"))

    books = query.all()
    # Ensure books are returned in the correct order
    ordered_books = [book.to_dict() for book in books]

    return jsonify(ordered_books)

# Retrieve a book by ID
@main_routes.route('/books/<int:id>', methods=['GET'])
def get_book(id):
    book = Book.query.get(id)
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404
    return jsonify(book.to_dict())

# Update a book by ID
@main_routes.route('/books/<int:id>', methods=['PUT'])
def update_book(id):
    data = request.get_json()   # Get JSON data from the request body
    
    book = Book.query.get(id)
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    # Validate genre if provided
    new_genre = data.get('genre')
    if new_genre:
        if new_genre not in VALID_GENRES:
            return jsonify({"error": "Invalid genre", "message": f"Genre must be one of the following: {', '.join(VALID_GENRES)}"}), 400
    
    # Update the book with provided data
    book.title = data.get('title', book.title)
    book.author = data.get('author', book.author)
    book.published_year = data.get('published_year', book.published_year)
    
    # Only update genre if it's valid
    if new_genre:
        book.genre = new_genre
    book.genre = data.get('genre', book.genre)

    db.session.commit()
    return jsonify({"message": "Book updated successfully"})

# Delete a book by ID
@main_routes.route('/books/<int:id>', methods=['DELETE'])
def delete_book(id):
    book = Book.query.get(id)
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    db.session.delete(book)
    db.session.commit()
    return jsonify({"message": "Book deleted successfully"})
