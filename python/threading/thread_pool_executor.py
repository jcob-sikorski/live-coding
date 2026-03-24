# The Square List: Create a function that squares a number.
# Use a ThreadPoolExecutor to square a list of numbers from 1 to 10 and print the results.

from concurrent.futures import ThreadPoolExecutor
import time

def square_number(num: int) -> None:
    print(f"Performing comuputation...")
    time.sleep(1)
    print(f"Result {num*num}")

with ThreadPoolExecutor(max_workers=3) as executor:
    # Submits tasks and returns a generator of results
    executor.map(square_number, [1, 2, 3, 4, 5])