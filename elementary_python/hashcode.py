def hash_code(s: str) -> str:
    if not s:
        return ""
    
    size = len(s)
    result = []

    for char in s:
        ascii_val = ord(char)

        new_val = (ascii_val + size) % 127

        if new_val < 33:
            new_val += 33
        
        result.append(chr(new_val))
        
    return "".join(result)

if __name__ == "__main__":
    print(hash_code("A"))
    print(hash_code("AB"))
    print(hash_code("BAC"))
    print(hash_code("Hello World"))
    print(ord('B'))
    print(chr(ord('A')))