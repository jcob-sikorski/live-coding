# Dictionary Transformation

# Given a dictionary prices = {'apple': 0.5, 'banana': 0.25, 'orange': 0.75},
# use map() and a lambda to create a new dictionary where all prices are
# increased by 20%.
# (Hint: You'll need to wrap the lambda output in a dict() constructor
# or use .items())

prices = {'apple': 0.5, 'banana': 0.25, 'orange': 0.75}

# Using round() is the cleanest way to handle the math here
increaser = (lambda x: (x[0], round(x[1] * 1.2, 2)))

new_prices = dict(map(increaser, prices.items()))

print(new_prices) 
# Output: {'apple': 0.6, 'banana': 0.3, 'orange': 0.9}