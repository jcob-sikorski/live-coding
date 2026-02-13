# The Task: Re-implement Python’s built-in open() as a context manager
# called opener. It must ensure the file is closed
# even if an error occurs inside the block.

from contextlib import contextmanager


@contextmanager
def opener(filename: str, mode: str = "r"):
    # 1. Setup: Open the resource
    f = open(filename, mode)
    try:
        # 2. Yield: Give the file object to the 'with' block
        yield f
    finally:
        # 3. Teardown: This is guaranteed to run,
        # even if an error occurs during the yield.
        print(f"Closing {filename}...")
        f.close()


try:
    with opener("timer.py", "w") as f:
        f.write("Learning context managers is fun!")
        raise ValueError("Something went wrong!")
except Exception as e:
    print(f"Caught error: {e}")
