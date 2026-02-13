# Exercise 1: The Inventory System (Dataclasses & Logic)
# Scenario: You are building a system for a warehouse. You need to store
# products, but you also need to calculate the total value of stock.

# Task: Create a @dataclass called Product with fields for name (str),
# price (float), and quantity (int).

# Bonus: Add a property or method to the dataclass that calculates the
# total_stock_value (price * quantity).

# Interview Focus: This tests if you know how to use modern Python features to
# avoid writing a manual __init__.

from dataclasses import dataclass

@dataclass
class Product:
    # These 'Type Hints' tell dataclass to build the __init__ for you
    name: str
    price: float
    quantity: int

    @property
    def total_stock_value(self) -> float:
        """Calculates the total value (Bonus task)"""
        return self.price * self.quantity

# Example Usage:
item = Product(name="Mechanical Keyboard", price=120.0, quantity=5)

print(f"Product: {item.name}")
print(f"Total Value: ${item.total_stock_value}")
# Notice we don't need () because of the @property decorator!