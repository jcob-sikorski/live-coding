# Exercise 1: The "Smart" Calculator (Level: Easy)
# The Goal: Build a function calculate(expression) that takes a tuple
# representing a mathematical operation and returns the result. This tests
# basic sequence matching and guards.

# Input: A tuple like ("add", 5, 10), ("sub", 10, 2), or ("mul", 3, 4).

# Requirements:

# Handle add, sub, and mul.

# Add a "guard" to prevent division by zero for a div operation.

# Use a wildcard to return an error message for unknown operations.

def calculator(expr):
    match expr:
        case ("add", x, y): return x + y
        case ("sub", x, y): return x - y
        case ("mul", x, y): return x * y
        case _: return "Unknown operation"
