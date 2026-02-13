import functools


def debug(func):
    @functools.wraps(func)  # Use the @ symbol here
    def wrapper(*args, **kwargs):
        print(f"--- Starting {func.__name__} ---")
        
        # Pass the arguments into the original function
        result = func(*args, **kwargs)
        
        print(f"--- Finished {func.__name__} ---")
        return result  # Return the actual result of the function
        
    return wrapper  # Return the function, don't call it!


@debug
def greet(name):
    print(f"Hello, {name}!")


# Now you can call it!
greet("Alice")