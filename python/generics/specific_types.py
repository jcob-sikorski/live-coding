# The Numeric Scaler
# Create a TypeVar constrained to only int or float. Write a function
# scale_numbers that takes a list of these numbers and a multiplier,
# returning a new list of the same type.

# from typing import TypeVar

# T = TypeVar("T", int, float)

# def scale_numbers(nums: list[T], multiplier: T) -> list[T]:
#     return [x * multiplier for x in nums]

def scale_numbers[T: (int, float)](nums: list[T], multiplier: T) -> list[T]:
    return [x * multiplier for x in nums]