# The Task: Write a manager SuppressErrors(*exceptions) that takes a list
# of exception types and prevents them from crashing
# the program if they occur inside the block.

from contextlib import contextmanager


@contextmanager
def suppress_errors(*exceptions):
    try:
        yield
    except exceptions:
        pass


# Testing for a single suppressed error
with suppress_errors(ZeroDivisionError):
    print("About to divide by zero...")
    result = 1 / 0
    print("This line will never run.")

print("The program continued successfully!")

# Testing for multiple suppressed errors
with suppress_errors(TypeError, ValueError):
    int("not_a_number")

# Testing an error that SHOULD crash (because it's not suppressed)
# This should still raise a KeyError
with suppress_errors(ZeroDivisionError):
    {}['invalid_key']
