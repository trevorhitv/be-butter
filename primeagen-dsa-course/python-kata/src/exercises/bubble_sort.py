
def bubble_sort(arr: list[int]):

    for i in range(len(arr)):
        for j in range(len(arr) - 1):
            if i > j:
                tmp = arr[i]
                arr[i] = arr[j]
                arr[j] = tmp

    return arr
