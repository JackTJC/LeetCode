package sort_alg

func HeapSort(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjust(arr, i, len(arr))
	}
	for i := len(arr) - 1; i >= 0; i-- {
		swap(arr, 0, i)
		adjust(arr, 0, i)
	}
}

func adjust(arr []int, pos, n int) {
	temp := arr[pos]
	k := 2*pos + 1
	for k < n {
		if k+1 < n && arr[k+1] < arr[k] {
			k++
		}
		if arr[k] < temp {
			// temp = arr[pos]
			arr[pos] = arr[k]
			// arr[k] = temp
			pos = k
		} else {
			break
		}
		k = 2*k + 1
	}
	arr[pos] = temp
}
func swap(arr []int, i, j int) {
	t := arr[i]
	arr[i] = arr[j]
	arr[j] = t
}
