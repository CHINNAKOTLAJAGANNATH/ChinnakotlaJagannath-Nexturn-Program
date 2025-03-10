# routes/transaction_routes.py
from flask import Blueprint, request, jsonify
from library_management_system import Library, BorrowLimitExceededException
from auth import require_auth

transaction_blueprint = Blueprint('transactions', __name__)
data_library = Library()

@transaction_blueprint.route('/borrow', methods=['POST'])
@require_auth
def borrow_book():
    data = request.get_json()
    member_id = data.get('member_id')
    book_id = data.get('book_id')
    try:
        data_library.lend_book(member_id, book_id)
        return jsonify({"message": f"Book ID {book_id} borrowed by Member ID {member_id}"}), 200
    except (ValueError, BorrowLimitExceededException) as e:
        return jsonify({"error": str(e)}), 400

@transaction_blueprint.route('/return', methods=['POST'])
@require_auth
def return_book():
    data = request.get_json()
    member_id = data.get('member_id')
    book_id = data.get('book_id')
    try:
        data_library.receive_return(member_id, book_id)
        return jsonify({"message": f"Book ID {book_id} returned by Member ID {member_id}"}), 200
    except ValueError as e:
        return jsonify({"error": str(e)}), 400
