# Library Management System

## Overview
This Library Management System is a Flask-based API designed to perform CRUD operations for managing books, members, and their borrowing activities. It adheres to Object-Oriented Programming (OOP) principles and includes features such as pagination, token-based authentication, and custom exception handling. The application maintains data locally in JSON files and is structured for scalability and ease of maintenance.

---

## How to Run the Project

1. **Install Python**:
   Ensure you have Python 3.x installed on your system.

2. **Clone the Repository**:
   Download or clone the repository containing the project files.

3. **Install Flask**:
   Install Flask by running the following command:
   ```bash
   pip install flask
   ```

4. **Run the Application**:
   Execute the `app.py` file to start the server:
   ```bash
   python app.py
   ```

   The application will run locally on `http://127.0.0.1:5000/`.

5. **Test the API**:
   Use `api_testing.py` to test all the operations:
   ```bash
   python api_testing.py
   ```
   or Use tools like curl, Postman, or a browser to interact with the API.

6. **Access API Endpoints**:
   - Books API: `http://127.0.0.1:5000/books`
   - Members API: `http://127.0.0.1:5000/members`
   - Transactions API: `http://127.0.0.1:5000/transactions`

---

## Design Choices

### 1. **File Structure**
The project is modularized for maintainability and scalability:
- `app.py`: Main application entry point.
- `routes/`: Contains Flask Blueprints for handling Books, Members, and Transactions endpoints.
- `library_management_system/`: Core classes (`Book`, `Member`, `Library`) and logic for managing the library.
- `data/`: Stores data persistently in JSON format (`books.json`, `members.json`).
- `api_testing.py`: Script to test API operations.

### 2. **OOP Principles**
- Encapsulation: Core logic is encapsulated within `Book`, `Member`, and `Library` classes.
- Inheritance: Members are categorized into Regular and Premium using shared attributes and unique borrowing limits.
- Polymorphism: Member borrowing limits are enforced through shared `borrow_book()` logic with class-specific constraints.

### 3. **Authentication**
- Token-based authentication is implemented using a decorator (`require_auth`) that checks for a valid token (`Bearer mysecrettoken`) in API requests.

### 4. **Pagination**
- Pagination is implemented for GET requests to `/books` and `/members` to handle large data efficiently.
- Query parameters `page` and `per_page` allow users to specify the results they want to retrieve.

### 5. **Error Handling**
- Built-in exceptions such as `ValueError` and `IndexError` are used for invalid inputs.
- A custom exception (`BorrowLimitExceededException`) is raised when a member exceeds their borrowing limit.

### 6. **Scalability**
- The current system uses JSON files for storage (`books.json`, `members.json`) but is designed to easily integrate SQL or NoSQL databases in the future.

---

## Assumptions and Limitations

### Assumptions
1. **Authentication**:
   - A static token (`Bearer mysecrettoken`) is used for simplicity.
   - In production, this would be replaced with a secure token management system (e.g., JWT).

2. **Borrowing Logic**:
   - Regular members can borrow up to 3 books.
   - Premium members can borrow up to 5 books.
   - Borrowing and returning books update their availability and member records.

3. **Data Storage**:
   - Data is stored locally in JSON format (`books.json`, `members.json`) for simplicity.
   - The system is designed to support database integration for scalability.

### Limitations
1. **Concurrency**:
   - JSON storage does not handle concurrent read/write operations efficiently.
   - A database should be used for concurrent access in a production environment.

2. **Authentication**:
   - Static token authentication is not secure for production use.
   - Token expiration, refresh mechanisms, and role-based access control are not implemented.

3. **Error Handling**:
   - Error messages are generic and may need to be localized or improved for user-friendliness.

4. **Data Validation**:
   - Basic validation is performed, but complex scenarios (e.g., invalid JSON structure) might need additional checks.

---

## Features

### CRUD Operations
- **Books**:
  - Add, update, delete, and fetch books.
  - Search books by title and/or author.
  - Example:
    ```bash
    GET /books/?title=Harry&author=Rowling&page=1&per_page=5
    ```

- **Members**:
  - Add, update, delete, and fetch members.
  - Categorized into Regular and Premium with specific borrowing limits.

### Borrow and Return Books
- Borrow books (if available and within the borrowing limit).
- Return books (and update their availability status).
- Example:
  ```bash
  POST /transactions/borrow
  {
      "member_id": 1,
      "book_id": 5
  }
  ```

### Search and Pagination
- Search books by title or author (or both).
- Pagination for books and members for efficient data retrieval.

### Authentication
- Token-based authentication required for all operations.

---

## Edge Cases Addressed

1. **Borrowing Unavailable Books**
   - The system ensures that a book’s availability is checked before processing a borrow operation.
   - If a book is already borrowed, an appropriate error message is returned, preventing duplicate borrowing attempts.

2. **Exceeding Borrowing Limit**
   - A custom exception, `BorrowLimitExceededException`, is raised when a member attempts to borrow more books than their membership type allows:
     - **Regular Members:** Limit of 3 books.
     - **Premium Members:** Limit of 5 books.

3. **Returning Books Not Borrowed**
   - The system validates whether the book being returned is present in the member’s borrowed list.
   - If a member attempts to return a book they haven’t borrowed, an error response is generated to prevent inconsistencies.

4. **Invalid Data Entry**
   - Input validation ensures that invalid data is rejected in the following cases:
     - Adding a book without specifying its title or author.
     - Registering a member without a valid name.
   - Meaningful error messages are returned for invalid input, helping users to correct their entries.

---

## Future Enhancements
1. **Database Integration**:
   Replace JSON storage with SQL (e.g., PostgreSQL) or NoSQL (e.g., MongoDB) databases.

2. **Advanced Authentication**:
   Implement JWT or OAuth for secure and scalable authentication.

3. **Frontend Integration**:
   Build a user-friendly frontend (e.g., React or Angular) to interact with the API.

4. **Improved Error Handling**:
   Add detailed error responses and localization support.

5. **Logging**:
   Add logging for debugging and monitoring API usage.

6. **Role-Based Access**:
   Introduce roles (e.g., Admin, Librarian, Member) for restricted access to certain operations.

