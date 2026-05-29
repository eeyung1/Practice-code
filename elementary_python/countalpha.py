def count_alpha(s: str) -> int:
    count = 0
    for char in s:
        if char.isalpha():
            count += 1
    return count

if __name__ == "__main__":
    print(count_alpha("Hello world"))
    print(count_alpha("H e l l o"))     
    print(count_alpha("H1e2l3l4o"))     
    print(count_alpha("123!@#"))        
    print(count_alpha(""))