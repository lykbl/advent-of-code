test = [1, 2, 3]


def do_smth(arr: list[int], i: int):
    if i > len(arr):
        return

    print(arr)
    arr[0] = None
    do_smth(arr, i + 1)


do_smth([1,2,3,4], 0)
