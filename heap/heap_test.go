package heap

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	data := []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	BuildMaxHeap(data)
	HeapSort(data)
	fmt.Printf("sorted array:%v", data)
}
