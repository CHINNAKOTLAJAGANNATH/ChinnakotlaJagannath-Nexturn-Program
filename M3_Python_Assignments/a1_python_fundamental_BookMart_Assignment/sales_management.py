
'''
3. Sales Management 
    o Sell a book to a customer (reduce the book quantity after a sale and log the transaction). 
    o View all sales records. 
'''

from customer_management import customers
from book_management import books, Book

class Transaction:
    def __init__(self, customer_name, book_title, quantity_sold):
        self.customer_name = customer_name
        self.book_title = book_title
        self.quantity_sold = quantity_sold

    def display_details(self):
        return f"Customer: {self.customer_name}, Book: {self.book_title}, Quantity Sold: {self.quantity_sold}"


sales = []

# Sell a book to a customer (reduce the book quantity after a sale and log the transaction). 
def sell_book(customer_name, book_title, quantity_sold):
    try:
        quantity_sold = int(quantity_sold)
        if quantity_sold <= 0:
            raise ValueError("Quantity must be positive.")
        
        book = next((b for b in books if b.title.lower() == book_title.lower()), None)
        if not book:
            return "Error: Book not found."
        if book.quantity < quantity_sold:
            return f"Error: Only {book.quantity} copies available. Sale cannot be completed."
        
        book.quantity -= quantity_sold
        sales.append(Transaction(customer_name, book.title, quantity_sold))
        return f"Sale successful! Remaining quantity: {book.quantity}"
    except ValueError as e:
        return f"Invalid input! {e}"

# View all sales records.
def view_sales():
    if not sales:
        return "No sales records available."
    return "\n".join(sale.display_details() for sale in sales)
