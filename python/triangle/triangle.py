def is_triangle(sides):
    if 0 in sides:
        return False
    if sides[0] + sides[1] < sides[2] or sides[0] + sides[2] < sides[1] or sides[1] + sides[2] < sides[0]:
        return False
    return True


def equilateral(sides):
    if not is_triangle(sides):
        return False
    return len(set(sides)) == 1


def isosceles(sides):
    if not is_triangle(sides):
        return False
    return sides[0] == sides[1] or sides[0] == sides[2] or sides[1] == sides[2]


def scalene(sides):
    if not is_triangle(sides):
        return False
    return len(set(sides)) == 3

