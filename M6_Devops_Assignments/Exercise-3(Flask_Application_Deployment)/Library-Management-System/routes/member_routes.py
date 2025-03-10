from flask import Blueprint, request, jsonify
from library_management_system import Member, Library
from auth import require_auth


member_blueprint = Blueprint('members', __name__)
data_library = Library()

@member_blueprint.route('/', methods=['GET'])
@require_auth
def get_members():
    page = int(request.args.get('page', 1))
    per_page = int(request.args.get('per_page', 5))

    start = (page - 1) * per_page
    end = start + per_page
    paginated_members = [
        {
            'member_id': member.member_id,
            'name': member.name,
            'membership': member.membership,
            'borrowed_books': [book.title for book in member.borrowed_books]
        }
        for member in data_library.members[start:end]
    ]

    return jsonify(paginated_members), 200

@member_blueprint.route('/', methods=['POST'])
@require_auth
def add_member():
    data = request.json
    try:
        new_member = Member(len(data_library.members) + 1, data['name'], data['membership'])
        data_library.register_member(new_member)
        return jsonify({
            'member_id': new_member.member_id,
            'name': new_member.name,
            'membership': new_member.membership
        }), 201
    except ValueError as e:
        return jsonify({'error': str(e)}), 400

@member_blueprint.route('/<int:member_id>', methods=['PUT'])
@require_auth
def update_member(member_id):
    data = request.json
    try:
        member = data_library.get_member(member_id)
        member.name = data.get('name', member.name)
        member.membership = data.get('membership', member.membership)
        data_library.save_members()
        return jsonify({
            'member_id': member.member_id,
            'name': member.name,
            'membership': member.membership
        }), 200
    except ValueError as e:
        return jsonify({'error': str(e)}), 404

@member_blueprint.route('/<int:member_id>', methods=['DELETE'])
@require_auth
def delete_member(member_id):
    try:
        member = data_library.get_member(member_id)
        data_library.members.remove(member)
        data_library.save_members()
        return '', 204
    except ValueError as e:
        return jsonify({'error': str(e)}), 404

