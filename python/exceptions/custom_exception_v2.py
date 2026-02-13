# The "Locked Account" Payload
# In security systems, we don't just say "Access Denied." We often tell the
# user why and how long they are blocked.

# The Challenge:

# Create a Custom Exception: Name it AccountLockedError.

# The Payload: It should take two arguments in __init__: username and
# minutes_remaining.

# The Logic: Write a function login(username, attempts). If attempts is
# greater than 3, raise your custom exception.

# The Catch: Wrap a call to login("admin", 5) in a try/except block.

# The Output: In the except block, use the exception's attributes to print:

# "Sorry admin, your account is locked. Try again in 15 minutes."

# Setup Code to Get You Started:
# Python
# class AccountLockedError(Exception):
#     # Your code here (Hint: don't forget super().__init__!)
#     pass

# def login(username, attempts):
#     # Your logic here
#     pass

# # Your try/except block here

class AccountLockedError(Exception):
    def __init__(self, username: str, minutes_remaining: int):
        self.username = username
        self.minutes_remaining = minutes_remaining
        # Make the message dynamic based on the username provided
        message = f"Account '{username}' is locked. Try again in {minutes_remaining} minutes."
        super().__init__(message)


def login(username, attempts):
    # Only raise the error if the business logic condition is met
    if attempts > 3:
        # For this exercise, let's assume 15 mins is the lockout time
        raise AccountLockedError(username, 15)
    print(f"Welcome back, {username}!")


# The Catch
try:
    login('Jacob', 5)
except AccountLockedError as e:
    # Accessing the payload attributes directly
    print(f"SECURITY ALERT: Sorry {e.username}, your account is locked. Try again in {e.minutes_remaining} minutes.")