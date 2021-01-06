def quickSort(arr):
    if len(arr) == 0 or len(arr) == 1:
        return arr
    else:
        return quickSort([i for i in arr[1:] if i < arr[0]]) + [arr[0]] + quickSort([i for i in arr[1:] if i >= arr[0]])


if __name__ == '__main__':
    print(quickSort([10, 9, 8, 7, 6, 5, 4, 3, 2, 1]))
    dict()
