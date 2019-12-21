package sort

func QuickSort(array []int32) {
	quickSort(array, 0, len(array)-1)
}

func quickSort(array []int32, left int, right int) {
	if left < right {
		partitionIndex := partition(array, left, right)
		// 不能包括partitionIndex
		quickSort(array, left, partitionIndex-1)  // partitionIndex 左边
		quickSort(array, partitionIndex+1, right) // partitionIndex 右边
	}
}

func partition(array []int32, left int, right int) int {
	p := left
	i := p + 1
	for j := i; j <= right; j++ {
		if array[j] < array[p] {
			array[j], array[i] = array[i], array[j]
			i++
		}
	}
	array[p], array[i-1] = array[i-1], array[p]
	return i - 1
}
