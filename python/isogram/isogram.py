def is_isogram(string):
    stripped_string = string.replace(" ", "").replace("-", "")
    string_set = set(''.join(stripped_string).lower())
    return len(string_set) == len(stripped_string)
