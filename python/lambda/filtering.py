# The Even/Odd Checker
# Write a lambda that returns True if a number is even and False if it is odd.
# Use this lambda with filter() to extract all even numbers from range(1, 21).

even_odd_checker = (lambda x: x % 2 == 0)

even_numbers = list(filter(even_odd_checker, range(1, 21)))

print(even_numbers)