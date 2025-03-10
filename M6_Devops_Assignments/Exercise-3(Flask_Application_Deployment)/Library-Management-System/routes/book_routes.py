from flask import Blueprint, request, jsonify
from library_management_system import Book, Library
from auth import require_auth
import json

book_blueprint = Blueprint('books', __name__)
data_library = Library()

@book_blueprint.route('/', methods=['GET'])
@require_auth
def get_books():
    search_title = request.args.get('title', '').lower()
    search_author = request.args.get('author', '').lower()
    page = int(request.args.get('page', 1))
    per_page = int(request.args.get('per_page', 5))

    # if search_title == '' and search_author == '':
    #     # Case 1: Retrieve all books without filters
    #     books = [book.__dict__ for book in data_library.book_collection]
    #     start = (page - 1) * per_page
    #     end = start + per_page
    #     paginated_books = books[start:end]
    #     return jsonify(paginated_books), 200

    # Case 2: Search for books by title or author
    filtered_books = [
        book.__dict__ for book in data_library.book_collection
        if (search_title in book.title.lower() if search_title else True) and
           (search_author in book.author.lower() if search_author else True)
    ]

    start = (page - 1) * per_page
    end = start + per_page
    paginated_filtered_books = filtered_books[start:end]

    return jsonify(paginated_filtered_books), 200

@book_blueprint.route('/', methods=['POST'])
@require_auth
def add_book():
    data = request.json
    try:
        new_book = Book(len(data_library.book_collection) + 1, data['title'], data['author'], data['category'])
        data_library.add_book(new_book)
        return jsonify(new_book.__dict__), 201
    except ValueError as e:
        return jsonify({'error': str(e)}), 400

@book_blueprint.route('/<int:book_id>', methods=['PUT'])
@require_auth
def update_book(book_id):
    data = request.json
    try:
        book = data_library.get_book(book_id)
        book.title = data.get('title', book.title)
        book.author = data.get('author', book.author)
        book.category = data.get('category', book.category)
        data_library.save_books()
        return jsonify(book.__dict__), 200
    except ValueError as e:
        return jsonify({'error': str(e)}), 404

@book_blueprint.route('/<int:book_id>', methods=['DELETE'])
@require_auth
def delete_book(book_id):
    try:
        book = data_library.get_book(book_id)
        data_library.book_collection.remove(book)
        data_library.save_books()
        return '', 204
    except ValueError as e:
        return jsonify({'error': str(e)}), 404

