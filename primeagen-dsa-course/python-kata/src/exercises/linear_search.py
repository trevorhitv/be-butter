
def linear_search(haystack: list[int], needle: int):
    for i in haystack:
        if needle == i:
            return True
    return False
