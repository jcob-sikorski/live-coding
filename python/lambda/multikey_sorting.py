# 5. Multi-Key Sorting
# You have a list of dictionaries representing products:
# products = [{'name': 'Laptop', 'price': 1200},
# {'name': 'Mouse', 'price': 25}, 'name': 'Monitor', 'price': 300}]
# Sort this list from highest price to lowest using a lambda as the key.

price_extractor = (lambda item: item['price'])

products = [
    {'name': 'Laptop', 'price': 1200},
    {'name': 'Mouse', 'price': 25},
    {'name': 'Monitor', 'price': 300}
]

# reverse=True goes inside the sorted() parentheses
sorted_products = sorted(products, key=price_extractor, reverse=True)

print(sorted_products)