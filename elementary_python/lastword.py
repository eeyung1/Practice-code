def last_word(s: str) -> str:
    if not s or s.isspace():
        return ""
    
    s = s.rstrip()

    last_space_index = s.rfind(" ")

    if last_space_index == -1:
        return s
    else:
        return s[last_space_index + 1:]

if __name__ == "__main__":
    print((last_word("this        ...       is sparta, then again, maybe    not")))
    print((last_word("lorem,ipsum")))
    print((last_word(" ")))
    print((last_word("singleword")))