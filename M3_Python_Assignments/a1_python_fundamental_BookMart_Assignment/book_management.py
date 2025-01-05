
''' 1. Book Management 
        o Add new books to the inventory (title, author, price, and quantity). 
        o View the list of all available books. 
        o Search for a book by its title or author. 
'''

class Book:
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def display_details(self):
        return f"Title: {self.title}, Author: {self.author}, Price: {self.price}, Quantity: {self.quantity}"


books = []

# Add new books to the inventory (title, author, price, and quantity).
def add_book(title, author, price, quantity):
    try:
        price = float(price)
        quantity = int(quantity)
        if price <= 0 or quantity <= 0:
            raise ValueError("Price and quantity must be positive.")
        books.append(Book(title, author, price, quantity))
        return "Book added successfully!"
    except ValueError as e:
        return f"Invalid input! {e}"

# View the list of all available books. 
def view_books():
    if not books:
        return "No books available."
    return "\n".join(book.display_details() for book in books)

# Search for a book by its title or author.
def search_book(query):
    results = [book.display_details() for book in books if query.lower() in book.title.lower() or query.lower() in book.author.lower()]
    return "\n".join(results) if results else "No books found."
