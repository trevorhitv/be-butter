from src.exercises.linear_search import linear_search

def test_linear_search():
    arr = [1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420]

    assert linear_search(arr, 69) == True
    assert linear_search(arr, 1336) == False
    assert linear_search(arr, 69420) == True
    assert linear_search(arr, 69421) == False
    assert linear_search(arr, 1) == True
    assert linear_search(arr, 0) == False
