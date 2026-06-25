def is_valid_camel_case(s: str) -> bool:
    if not s:
        return False
    if not s[0].isalpha():
        return False

    for i in range(1, len(s)):
        char = s[i] 
        if char.isdigit() or not char.isalnum():
            return False
        
        if char.isupper() and s[i-1].isupper():
            return False
    return True

def camel_to_snake_case(s: str) -> str:
    if not s:
        return ""

    if not is_valid_camel_case(s):
        return s
    
    result = []
    for i, char in enumerate(s):
        if char.isupper() and i > 0:
            result.append("_")
        result.append(char)

    return "".join(result)

if __name__ == "__main__":
    print(camel_to_snake_case("HelloWorld"))
    print(camel_to_snake_case("helloWorld"))
    print(camel_to_snake_case("camelCase"))
    print(camel_to_snake_case("CAMELtoSnackCASE"))
    print(camel_to_snake_case("camelToSnakeCase"))
    print(camel_to_snake_case("hey2"))
    print(camel_to_snake_case(""))
    print(camel_to_snake_case("Hello_World"))