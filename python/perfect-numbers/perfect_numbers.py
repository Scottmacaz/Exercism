def classify(number):
    if number <= 0:
        raise ValueError("number must be greater than zero.")
    
    s = sum(find_factors(number))
    
    if s == number:
        return "perfect"
    if s >= number:
        return "abundant"
    if s <= number:
        return "deficient"

def find_factors(number):
    factors = []
    for x in range(1, number):
        if number % x == 0:
            factors.append(x)
        x += 1
    return factors