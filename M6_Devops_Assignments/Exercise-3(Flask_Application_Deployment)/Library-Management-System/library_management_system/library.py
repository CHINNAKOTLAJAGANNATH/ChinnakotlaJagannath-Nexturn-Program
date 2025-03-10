from library_management_system.book import Book
from library_management_system.member import Member
import json


class Library:
    def __init__(self, books_file='data/books.json', members_file='data/members.json'):
        self.books_file = books_file
        self.members_file = members_file
        self.book_collection = self.load_books()
        self.members = self.load_members()


    def add_book(self, book):
        if book in self.book_collection:
            raise ValueError(f"Book with ID {book.book_id} already exists.")
        self.book_collection.append(book)
        self.save_books()

    def register_member(self, member):
        if member in self.members:
            raise ValueError(f"Member with ID {member.member_id} already exists.")
        self.members.append(member)
        self.save_members()

    def lend_book(self, member_id, book_id):
        member = self.get_member(member_id)
        book = self.get_book(book_id)
        member.borrow_book(book)
        self.save_books()
        self.save_members()

    def receive_return(self, member_id, book_id):
        member = self.get_member(member_id)
        book = self.get_book(book_id)
        member.return_book(book)
        self.save_books()
        self.save_members()

    def get_member(self, member_id):
        for member in self.members:
            if member.member_id == member_id:
                print("Member Found Sucessfully..")
                return member
            
        raise ValueError(f"Member with ID {member_id} does not exist.")  

    def get_book(self, book_id):
        for book in self.book_collection:
            if book.book_id == book_id:
                return book

        raise ValueError(f"Book with ID {book_id} does not exist.")
       
    def save_books(self):
        with open(self.books_file, 'w') as file:
            json.dump([obj.__dict__ for obj in self.book_collection], file, indent=4)

    def save_members(self):
        with open(self.members_file, 'w') as file:
            members_data = []
            for member in self.members:
                member_dict = member.__dict__.copy()
                member_dict['borrowed_books'] = [book.__dict__ for book in member.borrowed_books]
                members_data.append(member_dict)

            json.dump(members_data, file, indent=4)
            
    def load_books(self):
        try:
            with open(self.books_file, 'r') as file:
                data = json.load(file)
                books = []
                for book_data in data:
                    book = Book(book_data['book_id'], book_data['title'], book_data['author'], book_data['category'])
                    book.status = book_data.get('status', 'available')  # Set status after object creation
                    books.append(book)
                return books
                
        except FileNotFoundError as e:
            print(f"Error: {e}")
            return []
        except Exception as e:
            print(f"Error: {e}")
            return []

    def load_members(self):
        try:
            with open(self.members_file, 'r') as file:
                data = json.load(file)
                members = []
                for member_data in data:
                    member = Member(member_data['member_id'], member_data['name'], member_data['membership'])
                    # member.max_books = member_data.get('max_books', 3)
                    for book_data in member_data.get('borrowed_books', []):
                        book = self.get_book(book_data['book_id'])  
                        member.borrowed_books.append(book)
                    members.append(member)
                return members
        except FileNotFoundError as e:
            print(f"Error: {e}")
            return []
        except Exception as e:
            print(f"Error: {e}")
            return []
       
    

