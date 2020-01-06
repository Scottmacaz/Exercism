def square_of_sum(number):
    return sum([i for i in range(number+1)]) ** 2

def sum_of_squares(number):
    sum=0
    for i in range(number+1):
        sum += i ** 2
    return sum

def difference_of_squares(number):
    return square_of_sum(number) - sum_of_squares(number)
