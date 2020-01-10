def is_triangle(sides):
    a, b, c = sorted(sides)
    return min(sides) > 0 and a + b > c


def equilateral(sides):
    if not is_triangle(sides):
        return False
    return len(set(sides)) == 1


def isosceles(sides):
    if not is_triangle(sides):
        return False
    return len(set(sides)) <= 2


def scalene(sides):
    if not is_triangle(sides):
        return False
    return len(set(sides)) == 3

