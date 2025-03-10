import requests

BASE_URL = "http://127.0.0.1:5000"
HEADERS = {"Authorization": "you-will-never-guess"}

def main_menu():
    print("\nAPI Testing - Library Management System")
    print("1. Add a New Book")
    print("2. Update a Book")
    print("3. Delete a Book")
    print("4. Search Books by Title or Author")
    print("5. Get All Books")
    print("6. Register a New Member")
    print("7. Update a Member")
    print("8. Delete a Member")
    print("9. Get All Members")
    print("10. Borrow a Book")
    print("11. Return a Book")
    print("12. Exit")

def main():
    while True:
        main_menu()
        choice = input("Enter your choice: ")

        if choice == '1':  # Add Book
            title = input("Enter book title: ")
            author = input("Enter author name: ")
            category = input("Enter book category: ")
            response = requests.post(
                f"{BASE_URL}/books/",
                headers=HEADERS,
                json={"title": title, "author": author, "category": category}
            )
            print(response.json())

        elif choice == '2':  # Update Book
            book_id = input("Enter book ID to update: ")
            title = input("Enter new title (leave blank to keep current): ")
            author = input("Enter new author (leave blank to keep current): ")
            category = input("Enter new category (leave blank to keep current): ")
            data = {k: v for k, v in {"title": title, "author": author, "category": category}.items() if v}
            response = requests.put(
                f"{BASE_URL}/books/{book_id}",
                headers=HEADERS,
                json=data
            )
            print(response.json())

        elif choice == '3':  # Delete Book
            book_id = input("Enter book ID to delete: ")
            response = requests.delete(
                f"{BASE_URL}/books/{book_id}",
                headers=HEADERS
            )
            print(response.status_code, response.text or "Book deleted successfully.")

        elif choice == '4':  # Search Books by Title or Author
            title = input("Search by title (leave blank if not searching by title): ") 
            author = input("Search by author (leave blank if not searching by author): ") 
            response = requests.get(
                f"{BASE_URL}/books/",
                headers=HEADERS,
                params={"title": title, "author": author}
            )
            print(response.json())

        elif choice == '5':  # Get All Books
            page = input("Enter page number (leave blank for default): ") or 1
            per_page = input("Enter results per page (leave blank for default): ") or 5
            response = requests.get(
                f"{BASE_URL}/books/",
                headers=HEADERS,
                params={"page": int(page), "per_page": int(per_page)}
            )
            print(response.json())

        elif choice == '6':  # Register a Member
            name = input("Enter member name: ")
            membership = input("Enter membership type (Regular/Premium): ").capitalize()
            response = requests.post(
                f"{BASE_URL}/members/",
                headers=HEADERS,
                json={"name": name, "membership": membership}
            )
            print(response.json())

        elif choice == '7':  # Update Member
            member_id = input("Enter member ID to update: ")
            name = input("Enter new name (leave blank to keep current): ")
            membership = input("Enter new membership type (leave blank to keep current): ")
            data = {k: v for k, v in {"name": name, "membership": membership}.items() if v}
            response = requests.put(
                f"{BASE_URL}/members/{member_id}",
                headers=HEADERS,
                json=data
            )
            print(response.json())

        elif choice == '8':  # Delete Member
            member_id = input("Enter member ID to delete: ")
            response = requests.delete(
                f"{BASE_URL}/members/{member_id}",
                headers=HEADERS
            )
            print(response.status_code, response.text or "Member deleted successfully.")

        elif choice == '9':  # Get All Members
            page = input("Enter page number (leave blank for default): ") or 1
            per_page = input("Enter results per page (leave blank for default): ") or 5
            response = requests.get(
                f"{BASE_URL}/members/",
                headers=HEADERS,
                params={"page": int(page), "per_page": int(per_page)}
            )
            print(response.json())

        elif choice == '10':  # Borrow a Book
            member_id = input("Enter member ID: ")
            book_id = input("Enter book ID: ")
            response = requests.post(
                f"{BASE_URL}/transactions/borrow",
                headers=HEADERS,
                json={"member_id": int(member_id), "book_id": int(book_id)}
            )
            print(response.json())

        elif choice == '11':  # Return a Book
            member_id = input("Enter member ID: ")
            book_id = input("Enter book ID: ")
            response = requests.post(
                f"{BASE_URL}/transactions/return",
                headers=HEADERS,
                json={"member_id": int(member_id), "book_id": int(book_id)}
            )
            print(response.json())

        elif choice == '12':  # Exit
            print("\nExiting API Testing. Goodbye!")
            break

        else:
            print("\nInvalid choice. Please try again.")

if __name__ == "__main__":
    main()
