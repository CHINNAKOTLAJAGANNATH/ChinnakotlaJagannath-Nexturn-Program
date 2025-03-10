from flask import request, abort
from functools import wraps

def require_auth(func):
    @wraps(func)
    def decorated(*args, **kwargs):
        token = request.headers.get('Authorization')
        if not token or token != 'you-will-never-guess':
            abort(401, description='Unauthorized')
        return func(*args, **kwargs)
    return decorated