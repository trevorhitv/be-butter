from math import floor

def binary_search(haystack: list[int], needle: int):
    lo = 0
    hi = len(haystack)

    while lo < hi:
        m = floor(lo + (hi - lo) / 2)
        v = haystack[m]

        if v == needle:
            return True
        elif v > needle:
            hi = m
        else:
            lo = m + 1

    return False
