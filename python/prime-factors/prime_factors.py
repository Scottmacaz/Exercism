def factors(value):
    prime_factors = []
    current_divisor = 2

    while current_divisor <= value:
        if value % current_divisor == 0:
            prime_factors.append(current_divisor)
            value = int(value / current_divisor)
        else:
            current_divisor += 1

    return prime_factors
