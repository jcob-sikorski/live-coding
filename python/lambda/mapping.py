# Case Converter
# Create a lambda that takes a string and returns it in ALL CAPS.
# Then, use it inside a map() function to convert
# a list of lowercase names: ['alice', 'bob', 'charlie'].

# The lambda: just (lambda parameter: expression)
all_caps = (lambda a: a.upper())

names = ['alice', 'bob', 'charlie']

# Convert the map object to a list to view it
result = list(map(all_caps, names))

print(result)  # Output: ['ALICE', 'BOB', 'CHARLIE']