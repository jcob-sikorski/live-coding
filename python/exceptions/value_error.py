# Boundary Enforcement
# Write a function validate_percentage(n). If $n$ is less than 0 or greater
# than 100, raise a ValueError with a custom message.

def validate_percentage(n: int) -> bool:
    if not (0 <= n <= 100):
        # Using an f-string to show the user exactly what went wrong
        raise ValueError(f"Invalid percentage: {n}. Must be between 0 and 100.")

    return True

# If you run this, it stops the program immediately:
# validate_percentage(150)
