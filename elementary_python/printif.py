def printif(input):
    if len(input) <= 3 or len(input) == 0:
        return "G"
    return "Invalid Input\n"

print(printif("hel"))