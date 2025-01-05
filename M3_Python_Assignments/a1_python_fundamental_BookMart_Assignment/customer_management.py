
'''
2. Customer Management 
        o Add customer details (name, email, and phone number). 
        o View the list of all customers. 
'''

class Customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone

    def display_details(self):
        return f"Name: {self.name}, Email: {self.email}, Phone: {self.phone}"


customers = []

# Add customer details (name, email, and phone number).
def add_customer(name, email, phone):
    if not name or not email or not phone:
        return "All fields are required."
    customers.append(Customer(name, email, phone))
    return "Customer added successfully!"

# View the list of all customers. 
def view_customers():
    if not customers:
        return "No customers available."
    return "\n".join(customer.display_details() for customer in customers)
