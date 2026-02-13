# Input Sanitization
# Ask the user for their age using input().
# Use a try/except block to catch ValueError if the user
# types a string (like "twenty") instead of a number.
# Keep asking until they provide a valid integer.


while True:
    try:
        # 1. Take the input and immediately try to convert it
        age = int(input("What is your age? "))

        # 2. If the line above succeeds, we hit the break and leave the loop
        break

    except ValueError:
        # 3. If int() fails, we land here, print the message,
        # and the loop starts over from the top!
        print("Invalid input. Please enter a whole number (e.g., 25).")

print(f"Access granted. You are {age} years old.")