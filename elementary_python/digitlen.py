def digit_len(n: int, base: int) -> int:
    if base < 2 or base > 36:
        return -1
    
    if n == 0:
        return 1

    if n < 0:
        n = -n
    
    count = 0
    while n > 0:
        n = n//base
        count = count + 1
    
    return count

if __name__ == "__main__":
    print(digit_len(100, 10))      
    print(digit_len(100, 2))       
    print(digit_len(-100, 16))     
    print(digit_len(100, -1))      
    print(digit_len(0, 10))       
    print(digit_len(-45, 5))       