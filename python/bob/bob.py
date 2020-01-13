def response(hey_bob):
    hey_bob = hey_bob.strip()
    is_question = hey_bob[-1:] == '?'
    is_yelling = hey_bob.isupper()
    is_silent = not any(c.isalnum() for c in hey_bob)
    if is_silent and not is_question:
        return 'Fine. Be that way!'
    if is_yelling and is_question:
        return "Calm down, I know what I'm doing!"
    if is_question:
        return 'Sure.'
    if is_yelling:
        return 'Whoa, chill out!'

    return 'Whatever.'