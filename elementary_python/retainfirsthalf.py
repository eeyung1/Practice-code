def retainFirstHalf(input_str):
    if not input_str:
        return ""
    
    half = len(input_str) // 2

    if len(input_str) == 1:
        return input_str
    
    return input_str[:half]

if __name__ == "__main__":
    print(retainFirstHalf("Hello World"))
    print(retainFirstHalf("This is the 1st halfThis is the 2nd half"))
    print(retainFirstHalf("A"))
    print(retainFirstHalf(""))
