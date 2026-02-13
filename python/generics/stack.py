# List Wrapper
# Create a class Stack that can store a list of items of any single type $T$.
# Implement push(item: T) and pop() -> T.

class Stack[T]:
    def __init__(self):
        self.items: list[T] = []

    def push(self, item: T) -> None:
        self.items.append(item)

    def pop(self) -> T:
        return self.items.pop()
