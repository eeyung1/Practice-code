def has_uppercase(arg: str) -> bool:
    for char in arg:
        if char.isupper():
            return True
    return False

if __name__ == "__main__":
    print(has_uppercase("hello"))      #
    print(has_uppercase("Hello"))      
    print(has_uppercase("PYTHON"))     
    print(has_uppercase("123abc"))     
    print(has_uppercase("abcDef"))     
    print(has_uppercase(""))