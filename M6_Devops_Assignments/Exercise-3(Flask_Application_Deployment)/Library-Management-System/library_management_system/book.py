
class Book:
    def __init__(self, book_id, title, author, category):
        if not title or not author:
            raise ValueError("Title and Author cannot be empty.")
        self.book_id = book_id
        self.title = title
        self.author = author
        self.category = category
        self.status = "available"

    def borrow(self):
        if self.status == "borrowed":
            raise ValueError(f"{self.title} is currently Not Available..")
        self.status = "borrowed"

    def return_book(self):
        self.status = 'available'

    def __str__(self):
        return f"Book({self.book_id}): '{self.title}' by {self.author} - Status: {self.status}"

        
