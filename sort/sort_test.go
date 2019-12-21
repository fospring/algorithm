package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := []int32{2, 8, 7, 1, 3, 5, 6, 4}
	QuickSort(data)
	fmt.Printf("sorted array:%v", data)
}
