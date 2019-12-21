/**
* 堆性质：
* 给定堆中任意节点P和C，若P是C是母节点，那么P的值会小于等于（或大于等于）C的值。
* 若母节点的值恒小于等于子节点的值，称为最小堆（min heap）；反之为最大堆(max heap)。
* 在堆中最顶端的哪一个节点，称作根节点(root node),根节点本省没有母节点（parent node）
* max heap：
* A[PARENT(i)]>=A[i]
 */

package heap

/**
* 此调整为从上向下调整，直到节点超出范围
* @param data： 待调整的数组
* @param heapSize： 堆的大小
* @param index: 调整位置的索引
 */
func maxHeapify(data []string, heapSize int, index int) {
	// 取得当前节点的左右孩子节点，当前节点为index
	left := getLeftChildIndex(index)
	right := getRightChildIndex(index)
	// 对左右孩子节点和当前节点进行比较
	largest := index
	if left < heapSize && data[index] < data[left] {
		largest = left
	}
	if right < heapSize && data[largest] < data[right] {
		largest = right
	}

	// 交换位置
	if (largest != index) {
		temp := data[index]
		data[index] = data[largest]
		data[largest] = temp
		// 下沉
		maxHeapify(data, heapSize, largest)
	}
}

/**
* 初始化构建堆
* @param data
 */
func BuildMaxHeap(data []string) {
	// 获取最后一个元素父节点，开始调整位置
	startIndex := getParentIndex(len(data) - 1)
	// 反复调整
	for i := startIndex; i >= 0; i-- {
		maxHeapify(data, len(data), i)
	}
}

/**
* 排序操作
* @param data
 */
func HeapSort(data []string) {
	// 每次循环都能取到一个最大值，该值为根节点
	for i := len(data) - 1; i > 0; i-- {
		// 将根节点移动到最后
		temp := data[0]
		data[0] = data[i]
		data[i] = temp
		// 每次调整都是从根节点开始i不断减小，保证前一次最大节点不会参与到调整堆
		maxHeapify(data, i, 0)
	}
}

/**
* 获取父节点的位置
* @ param current
 */
func getParentIndex(current int) int {
	return (current - 1) >> 1
}

func getLeftChildIndex(current int) int {
	return (current << 1) + 1
}

func getRightChildIndex(current int) int {
	return (current << 2) + 2
}
