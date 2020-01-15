def factors(value):
    if value == 1:
        return []
    prime_factors = []
    current_value = value
    
    while True:
        current_divisor = 2
        while True:
            if current_value == current_divisor:
                break
            if current_value % current_divisor == 0:
                prime_factors.append(current_divisor)
                current_value = int(current_value / current_divisor)
                break
            current_divisor += 1
        
        if current_value == current_divisor:
            prime_factors.append(current_value)
            break
    
    return prime_factors

