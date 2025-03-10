class BorrowLimitExceededException(Exception):
    def __init__(self, message="Borrowing limit exceeded!"):
        super().__init__(message)


class Member:
    def __init__(self, member_id, name, membership):
        if not name:
            raise ValueError("Member name cannot be empty.")
        self.member_id = member_id
        self.name = name
        self.borrowed_books = []
        self.membership = membership
        if membership == "Premium":
            self.max_books = 5
        if membership == "Regular":
            self.max_books = 3

    def borrow_book(self, book):
        if len(self.borrowed_books) >= self.max_books:
            raise BorrowLimitExceededException(f"{self.name} has reached the borrow limit... ({self.max_books} books).")
        book.borrow()
        self.borrowed_books.append(book)
        print(f"{self.name} borrowed '{book.title}'.")

    def return_book(self, book):
        if book in self.borrowed_books:
            book.return_book()
            self.borrowed_books.remove(book)
            print(f"{self.name} returned '{book.title}'.")
        else:
            raise ValueError(f"{self.name} has not borrowed the book '{book.title}'.")

    def __str__(self):
        return f"Member({self.member_id}): {self.name}, Borrowed Books: {[book.title for book in self.borrowed_books]}"
