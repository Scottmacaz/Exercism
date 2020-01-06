def is_valid(isbn):
    isbn = isbn.replace("-", "")
    if len(isbn) != 10:
        return False
    isbnlist = list(isbn)
    if not isbnlist[9].isnumeric() and isbnlist[9] != 'X':
        return False
    if isbnlist[9] == 'X':
        isbnlist[9] = 10
    isbnsum = 0
    isbnmultiplier = 10
    for d in isbnlist:
        if not str(d).isnumeric():
            return False
        isbnsum += int(d) * isbnmultiplier
        isbnmultiplier -= 1

    return isbnsum % 11 == 0
