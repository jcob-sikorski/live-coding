# The First & Last Pair
# Write a function get_edges that takes a list of items of type $T$ and
# returns a tuple containing the first and last elements. Handle the type
# hinting so that if a list[int] is passed,the IDE knows
# the result is tuple[int, int].

def get_edges[T](items: list[T]) -> tuple[T, T]:
    return (items[0], items[-1])
