# 3. The Execution Timer
# Task: Create a decorator @timer that measures how many seconds a function
# takes to execute and prints the duration.

import functools
import time


def timer(func):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        start_time = time.perf_counter()

        try:
            result = func(*args, **kwargs)
            return result
        finally:
            end_time = time.perf_counter()
            duration = end_time - start_time
            # Using .4f ensures we see 0.0001 instead of 1e-04
            print(f"DEBUG: '{func.__name__}' executed in {duration:.4f}s")

    return wrapper

@timer
def fast_worker():
    """Does absolutely nothing, very quickly."""
    pass


fast_worker()