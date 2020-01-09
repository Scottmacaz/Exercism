def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("strands must be equal length")

    differences = 0
    for pos, a in enumerate(strand_a):
        if a != strand_b[pos]:
            differences += 1
    return differences
