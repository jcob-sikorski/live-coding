# Sorting by Last Name
# Given a list of full names:
# ['James Bond', 'Indiana Jones', 'Han Solo', 'Luke Skywalker'].
# Use sorted() and a lambda to sort this list alphabetically by the last name.

surname = (lambda x: x.split()[-1])

people = ['James Bond', 'Indiana Jones', 'Han Solo', 'Luke Skywalker']

sorted(people, key=surname)