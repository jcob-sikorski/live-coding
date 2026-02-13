# The Task: Use the @contextlib.contextmanager decorator to create a timer.
# It should record the start time, yield, and then print the elapsed time.

from contextlib import contextmanager
import time


@contextmanager
def timer(label):
    start = time.perf_counter()

    try:
        yield
    finally:
        end = time.perf_counter()
        elapsed = end - start
        print(f"[{label}] Elapsed time: {elapsed:.4f} seconds")


with timer("Counting to a million"):
    total = sum(range(1_000_000))
    print(f"Done! Sum is {total}")