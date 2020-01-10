def get_products(min_factor, max_factor):
    products = set()
    for item1 in range(min_factor, max_factor+1):
        for item2 in range(min_factor, max_factor+1):
            products.add(item1 * item2)
    return list(products)


def find_palindromes(num_list):
    palindromes = []
    for n in num_list:
        if str(n) == "".join(reversed(str(n))):
            palindromes.append(n)
    print("done")
    return palindromes

def find_factors(num, starting_num, ending_num ) :
    factors = []
    for x in range(starting_num, ending_num + 1):
        for y in range(starting_num, ending_num + 1):
            if x * y == num :
                if not [y,x] in factors: 
                    factors.append([x,y])
    print ("done.")
    return factors

def largest(min_factor, max_factor):
    if min_factor > max_factor:
        raise ValueError("min_factor must be greater than max_factor")
    palindromes = find_palindromes(get_products(min_factor, max_factor))
    if not palindromes:
        return None, []
    largest_num = max(palindromes)
    return largest_num, find_factors(largest_num, min_factor, max_factor)


def smallest(min_factor, max_factor):
    if min_factor > max_factor:
        raise ValueError("min_factor must be greater than max_factor")
    palindromes = find_palindromes(get_products(min_factor, max_factor))
    if not palindromes:
        return None, []
    smallest_num = min(palindromes)
    return smallest_num, find_factors(smallest_num, min_factor, max_factor)


print (find_factors(9, 1, 9))