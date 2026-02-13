# 2. The Logger (With Arguments)
# Task: Create a decorator @log_args that prints the name of the function being
# called and the arguments (*args, **kwargs) passed to it.

import functools


def log_args(func):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        # 1. Print the logs as requested
        print(f"-- Function name: {func.__name__}")
        print(f"-- Arguments called: args={args} kwargs={kwargs}")

        # 2. IMPORTANT: Capture the result and return it
        result = func(*args, **kwargs)
        return result

    return wrapper


@log_args
def worker(*tasks):  # Changed **args to *tasks to accept multiple positional strings
    print(f"Working on: {', '.join(tasks)}")


# Now this call works perfectly!
worker("cooking", "cleaning", "sleeping")
