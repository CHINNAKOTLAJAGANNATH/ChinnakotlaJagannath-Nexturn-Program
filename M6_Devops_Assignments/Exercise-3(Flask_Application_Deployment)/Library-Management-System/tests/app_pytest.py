import pytest
from app import app

# Fixtures for setting up the Flask test client
@pytest.fixture
def client():
    """Creates a test client for the Flask app."""
    app.config["TESTING"] = True
    client = app.test_client()
    return client

# Fixtures for Sample Data
@pytest.fixture
def new_book():
    return {"title": "Python Basics", "author": "John Doe", "category": "Programming"}

@pytest.fixture
def updated_book():
    return {"title": "Advanced Python"}

@pytest.fixture
def new_member():
    return {"name": "Alice", "membership": "Regular"}

@pytest.fixture
def updated_member():
    return {"name": "Alice Johnson", "membership": "Premium"}

@pytest.fixture
def borrow_data():
    return {"member_id": 1, "book_id": 1}

@pytest.fixture
def return_data():
    return {"member_id": 1, "book_id": 1}

HEADERS = {"Authorization": "you-will-never-guess"}


# API Tests Using Flask Test Client

def test_add_book(client, new_book):
    response = client.post("/books/", headers=HEADERS, json=new_book)
    assert response.status_code == 201
    assert response.json["title"] == new_book["title"]

def test_update_book(client, updated_book):
    book_id = 1  # Assuming book with ID 1 exists
    response = client.put(f"/books/{book_id}", headers=HEADERS, json=updated_book)
    assert response.status_code == 200
    assert response.json["title"] == updated_book["title"]

def test_delete_book(client):
    book_id = 1  # Assuming book with ID 1 exists
    response = client.delete(f"/books/{book_id}", headers=HEADERS)
    assert response.status_code == 204

def test_search_books(client):
    query = "Python"
    response = client.get(f"/books/?title={query}", headers=HEADERS)
    assert response.status_code == 200
    assert isinstance(response.json, list)

def test_get_all_books(client):
    response = client.get("/books/", headers=HEADERS)
    assert response.status_code == 200
    assert isinstance(response.json, list)

def test_register_member(client, new_member):
    response = client.post("/members/", headers=HEADERS, json=new_member)
    assert response.status_code == 201
    assert response.json["name"] == new_member["name"]

def test_update_member(client, updated_member):
    member_id = 1
    response = client.put(f"/members/{member_id}", headers=HEADERS, json=updated_member)
    assert response.status_code == 200
    assert response.json["name"] == updated_member["name"]

def test_delete_member(client):
    member_id = 1
    response = client.delete(f"/members/{member_id}", headers=HEADERS)
    assert response.status_code == 204

def test_get_all_members(client):
    response = client.get("/members/", headers=HEADERS)
    assert response.status_code == 200
    assert isinstance(response.json, list)

def test_borrow_book(client, borrow_data):
    response = client.post("/transactions/borrow", headers=HEADERS, json=borrow_data)
    assert response.status_code in [200, 400]  # Allow for book already borrowed

def test_return_book(client, return_data):
    response = client.post("/transactions/return", headers=HEADERS, json=return_data)
    assert response.status_code == 200
