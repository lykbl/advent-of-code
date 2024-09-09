def binary_search(tails, target, high):
    low = 0
    while low < high:
        mid = (low + high) // 2
        if tails[mid] < target:
            low = mid + 1
        else:
            high = mid
    return low

def longest_increasing_subsequence(arr):
    if not arr:
        return 0
    
    n = len(arr)
    tails = [0] * n
    size = 0
    
    for x in arr:
        i = binary_search(tails, x, size)
        tails[i] = x
        print(f"Current: {x}, tail is: {i}, tails: {tails}")
        size = max(size, i + 1)

    print(tails)
    
    return size

# Example usage
# arr = [1, 7, 4, 8, 3, 6, 2, 5]
arr = [2, 6, 5, 1, 7, 4, 8, 3]
print(f"Length of LIS: {longest_increasing_subsequence(arr)}")
