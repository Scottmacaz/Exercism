def classify(number):
    if number <= 0:
        raise Exception("number must be greater than zero.")
    
    s = sum(find_factors(number))
    print(s)

    if s == number:
        return "perfect"
    if s >= number:
        return "abundant"
    if s <= number:
        return "deficient"


def find_factors(number):
    factors = []
    x = 1
    while x < number:
        if (number % x == 0):
            factors.append(x)
        x += 1
    return factors


# print(classify(0))
