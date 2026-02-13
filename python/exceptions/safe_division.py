# The "Safe Division"
# Write a function safe_divide(a, b) that divides $a$ by $b$. If $b$ is zero,
# catch the ZeroDivisionError and return None instead of crashing.

from typing import Optional


def safe_divide(a: float, b: float) -> Optional(int):
    try:
        return a / b
    except ZeroDivisionError:
        return None


# Test it out
print(safe_divide(10, 2))  # Output: 5.0
print(safe_divide(10, 0))  # Output: None
