# The Challenge: "The Execution Timer"
# Question: "Write a decorator called @timer that calculates how long a
# function takes to execute.
# It should print the function's name and the duration in seconds.
# For bonus points, ensure the decorated function doesn't lose its identity
# (metadata) and can handle any number of arguments."

import time
import functools


def timer(func):
    @functools.wraps(func)  # Crucial: Preserves __name__ and docstrings
    def wrapper(*args, **kwargs):
        start_time = time.perf_counter()  # More precise than time.time()

        result = func(*args, **kwargs)   # Call the actual function

        end_time = time.perf_counter()
        duration = end_time - start_time

        print(f"Function {func.__name__!r} took {duration:.4f}s")
        return result
    return wrapper


# Usage:
@timer
def heavy_computation(n):
    """Simulates a heavy task."""
    return sum(i**2 for i in range(n))


heavy_computation(1000000)