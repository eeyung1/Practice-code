def count_digits(arg: str) -> int:
    count = 0
    for char in arg:
        if char.isdigit():
            count += 1
    return count

if __name__ == "__main__":
    print(count_digits("Hello"))
    print(count_digits("Hello123"))
    print(count_digits("abc1d2e3f"))
    print(count_digits("1234567890"))
    print(count_digits("No numbers here!"))