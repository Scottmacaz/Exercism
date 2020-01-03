def slices(series, length):
    if length < 1:
        raise ValueError("Length must be a positive integer.")
    if len(series) < length:
        raise ValueError("series length must be longer then length parameter")

    rsp = []
    for i in range(len(series)-length+1):
        rsp.append(series[i: i+length])
    return rsp

