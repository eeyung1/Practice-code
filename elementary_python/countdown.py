def countdown(i):
    print(i)

    if i <= 0: 
        return
    countdown(i-1)

countdown(7)