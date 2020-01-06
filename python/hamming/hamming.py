def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("strands must be equal length")

    position = 0
    differences = 0
    for a in strand_a:
        if a != strand_b[position]:
            differences += 1
        position += 1

    return differences
