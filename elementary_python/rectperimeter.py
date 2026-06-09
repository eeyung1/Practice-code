def rectperimeter(w, h):
    if w < 0 or h < 0:
        return -1
    return 2 * (w + h)

if __name__ == "__main__":
    print(rectperimeter(2, 4))
    print(rectperimeter(-1, 4))