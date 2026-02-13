# Custom Exceptions
# Exercise 7: The "Coffee Machine" Error
# Create a custom exception class called InsufficientWaterError. Write a class
# CoffeeMachine with a method brew(). If a water_level variable is below 10,
# raise your custom exception.

class InsufficientWaterError(Exception):
    """Raised when the water level is too low to brew coffee."""
    def __init__(self, current_level: int, required_level: int = 10):
        self.current_level = current_level
        self.required_level = required_level
        # Pass a clean message to the base Exception class
        message = f"Insufficient water: {current_level}ml (Need at least {required_level}ml)"
        super().__init__(message)


class CoffeeMachine:
    def brew(self, water_level: int):
        if water_level < 10:
            # We pass the data into our custom exception
            raise InsufficientWaterError(water_level)
        print("☕ Brewing a perfect cup of coffee!")


# Testing the machine
machine = CoffeeMachine()
try:
    machine.brew(5)
except InsufficientWaterError as e:
    print(f"Alert: {e}")
    print(f"Please add {e.required_level - e.current_level}ml more water.")