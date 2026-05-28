def check_number(arg: str) -> bool:

    for char in arg:
        if char.isdigit():
            return True
    return False

if __name__ == "__main__":
    print(check_number("Hello"))
    print(check_number("Hello"))
    print(check_number("aby123"))
    print(check_number("python"))