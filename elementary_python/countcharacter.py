def count_char(text: str, char: str) -> int:
    if not text:
        return 0
    
    count = 0
    for c in text:
        if c == char:
            count += 1
    return count

if __name__ == "__main__":
    print(count_char("Hello world", "l"))
    print(count_char("5 balloons", "5"))            
    print(count_char("   ", " "))                   
    print(count_char("The 7 deadly sins", "7"))     
    print(count_char("Hello", "x"))                 
    print(count_char("", "a"))                     