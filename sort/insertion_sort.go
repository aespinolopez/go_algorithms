package sort

func InsertionSort(arr []int) {
	for j := 2; j < len(arr); j++ {
		var key = arr[j]
		var i = j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i = i - 1
		}
		arr[i+1] = key
	}
}
