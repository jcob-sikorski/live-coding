# The List Indexer
# Write a function that takes a list and an index. Try to return the element
# at that index. Handle two specific errors: IndexError (if the index is out
# of range) and TypeError (if the index provided isn't an integer).

from typing import TypeVar, Any

T = TypeVar('T')

def get_item(elements: list[T], index: Any) -> T | None:
    try:
        return elements[index]
    except IndexError:
        print(f"Error: Index {index} is out of range.")
    except TypeError:
        print(f"Error: '{index}' is a {type(index).__name__}, but an integer is required.")
    return None


# Test cases
my_list = ["Apple", "Banana", "Cherry"]

get_item(my_list, 10)     # Triggers IndexError
get_item(my_list, "one")  # Triggers TypeError
