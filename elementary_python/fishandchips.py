def fishandchips(num):
    if num < 0:
        return "error: number is negative"
    if num % 2 == 0 and num % 3 == 0:
        return "fish and chips"
    elif num % 2 == 0:
        return "fish"
    elif num % 3 == 0:
        return "chips"
    else:
        return "error: non divisible"

if __name__ == "__main__":
    print(fishandchips(4))
    print(fishandchips(9))
    print(fishandchips(6))
    print(fishandchips(-6))