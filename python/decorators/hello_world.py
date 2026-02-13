# 1. The "Hello World" Decorator
# Task: Create a decorator called @start_stop that prints "Starting..." before
# a function runs and "Finished!" after it completes.

import functools


def start_stop(func):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        print("Starting...")

        func(*args, **kwargs)

        print("Finished!")

    return wrapper


@start_stop
def worker():
    print("working")


worker()
