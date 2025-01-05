
from app import create_app
from app.models import db

app = create_app()

with app.app_context():
    db.create_all()         # Creates all tables defined in the models
    # Optional: Add sample data
    # sample_books = [
    #     Book(title="The Great Gatsby", author="F. Scott Fitzgerald", published_year=1925, genre="Fiction"),
    #     Book(title="To Kill a Mockingbird", author="Harper Lee", published_year=1960, genre="Fiction"),
    #     Book(title="Dune", author="Frank Herbert", published_year=1965, genre="Sci-Fi")
    # ]
    # db.session.bulk_save_objects(sample_books)

    # db.session.commit()
    print("Database initialized successfully!")

