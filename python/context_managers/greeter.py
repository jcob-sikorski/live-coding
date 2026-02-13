# The Task: Create a class Greeter that prints "Hello!"
# when entering the with block and "Goodbye!" when exiting.

class Greeter:
    def __enter__(self) -> Greeter:
        print("Hello!")
        return self

    def __exit__(self, exc_type, exc_val, exc_tb) -> None:
        print("Goodbye!")


with Greeter():
    print("I'm hanging out inside the block.")