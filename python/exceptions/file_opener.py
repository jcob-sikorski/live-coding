# The Full Lifecycle
# Create a script that:

# Opens a file in try.

# Processes the data in else (if the file opened successfully).

# Always closes the file in finally, regardless of whether an error occurred.

file = None  # 1. Initialize to None so finally can check it
try:
    # 2. Only the "risky" opening happens here
    file = open('example.txt', 'r')
except FileNotFoundError:
    print("Error: The file does not exist.")
except Exception as e:
    print(f"An unexpected error occurred: {e}")
else:
    # 3. This runs ONLY if the try block succeeded
    content = file.read()
    print(f"Success! File has {len(content)} characters.")
finally:
    # 4. This runs NO MATTER WHAT
    if file:
        file.close()
        print("File handle closed safely.")
    else:
        print("No file was opened, nothing to close.")
