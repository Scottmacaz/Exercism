def is_armstrong_number(number):
    e = len(str(number))
    sum = 0
    for d in str(number):
        sum += int(d) ** e
    return sum == number
