def steps(number):
    if number <= 0:
        raise ValueError("number must be greater than zero")

    i = 0
    while number > 1:
        number = conjencture(number)
        i += 1
    return i

def conjencture(num):
    if num % 2 == 0:
        return num / 2
    
    return num * 3 + 1



#print(steps(12))