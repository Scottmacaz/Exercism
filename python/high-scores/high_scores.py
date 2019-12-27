def latest(scores):
    return scores.pop()


def personal_best(scores):
    scores.sort()
    return scores.pop()


def personal_top_three(scores):
    scores.sort(reverse=True)
    return scores[:3]
