def find_anagrams(word, candidates):
    return [candidate for candidate in candidates if len(word) == len(candidate) and word.lower() != candidate.lower() and sorted(word.lower()) == sorted(candidate.lower())]
