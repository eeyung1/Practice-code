def first_word(input):
    if not input:
        return ""

    result = []
    started = False

    for char in input:
        if char == " ":
            if started:
                break
            continue
        
        started = True
        result.append(char)
    
    return "".join(result)

if __name__ == "__main__":
    print(first_word("hello there"))
    print(first_word(" hi there"))
    print(first_word(""))
    print(first_word("hello   .........  bye"))
    print(first_word("singleword"))