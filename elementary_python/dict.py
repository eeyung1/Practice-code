book = dict ()
print("input a value to search items")
good = input()

book["apple"] = 0.67
book["milk"] = 1.49
book["avocado"] = 1.49

value = book.get(good)

# print(book)
print(value)
# print(good)