package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := NewHeap([]int{3, 1, 5, 3, 8, 3}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	heap.HeapInit()
	heap.Print()
	fmt.Println(heap.HeapPop(), heap.HeapPop(), heap.HeapPop())
	heap.HeapPush(4)
	heap.HeapPush(6)
	heap.Print()
	heap.HeapSort()
	heap.Print()
}
