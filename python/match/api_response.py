# Exercise 2: API Response Parser (Level: Medium)
# The Goal: In a real-world scenario, APIs often return nested dictionaries
# with varying structures based on success or failure. This tests nested
# dictionary matching and type checking.

# Input: A dictionary.

# Success: {"status": "success", "data": {"id": 1, "value": 100}}

# Error: {"status": "error", "error_code": 404, "msg": "Not Found"}

# Requirements:

# If success, return only the value.

# If error, return a formatted string: "Error {code}: {msg}".

# If the input is any other dictionary, return "Malformed Response".


def response_parser(response):
    match response:
        case {"status": "success", "data": {"value": v}}:
            return v
        case {"status": "error", "error_code": code, "msg": m}:
            return f"Error {code}: {m}"
        case _:
            return "Malformed Response"
