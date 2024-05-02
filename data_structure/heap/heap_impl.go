package heap

import "fmt"

type Heap struct {
	// index 存放数据的索引,data[index[i]]是真实的数据
	index []int
	data  any
}

func NewHeap(index []int, value any) *Heap {
	return &Heap{
		index: index,
		data:  value,
	}
}

func (h *Heap) Print() {
	for i := 0; i < h.Len(); i++ {
		fmt.Print(h.GetValue(i), " ")
	}
	fmt.Println()
}

func (h *Heap) Len() int {
	return len(h.index)
}

func (h *Heap) HeapInit() {
	for i := h.Len()/2 - 1; i >= 0; i-- {
		h.heapDown(i, h.Len()-1)
	}
}

func (h *Heap) HeapSort() {
	for i := h.Len() - 1; i > 0; i-- {
		h.swapIndex(0, i)
		h.heapDown(0, i-1)
	}
}

// HeapPop 返回堆顶的元素下标
func (h *Heap) HeapPop() int {
	if h.Len() == 0 {
		return -1
	}
	x := h.GetValue(0)
	h.swapIndex(0, h.Len()-1)
	h.index = h.index[:h.Len()-1]
	h.heapDown(0, h.Len()-1)
	return x
}

// HeapPush push一个元素到index中
func (h *Heap) HeapPush(i int) {
	h.index = append(h.index, i)
	h.heapUp()
}

// heapUp 数据上浮操作
func (h *Heap) heapUp() {
	// 默认需要up的元素位于index末尾,因为插入通常是append到最后
	curIndex := h.Len() - 1
	newValue := h.GetValue(curIndex)
	parent := (curIndex - 1) / 2
	for parent >= 0 && newValue > h.GetValue(parent) {
		h.swapIndex(parent, curIndex)
		curIndex = parent
		parent = (curIndex - 1) / 2
	}
}

// heapDown 数据下沉操作
func (h *Heap) heapDown(parent, end int) {
	// heap的定义,i的子节点分别为2*i+1和2*i+2
	child := 2*parent + 1
	for child <= end { // 是否取等号需要根据end的定义,如果是len(index)-1,就能取等号
		if child+1 <= end && h.GetValue(child) < h.GetValue(child+1) {
			child++
		}
		if h.GetValue(parent) < h.GetValue(child) {
			h.swapIndex(parent, child)
			parent = child
			child = 2*parent + 1
		} else {
			break
		}
	}
}

// swapIndex 交换index中下标为i和j的值
func (h *Heap) swapIndex(i, j int) {
	h.index[i], h.index[j] = h.index[j], h.index[i]
}

// GetValue 根据index中的下标i获取对应的数值
func (h *Heap) GetValue(i int) int {
	switch x := h.data.(type) {
	case []int:
		return x[h.index[i]]
	case map[int]int:
		return x[h.index[i]]
	}

	return -1
}

// GetIndex 返回堆顶元素,默认为index
func (h *Heap) GetIndex(i int) int {
	return h.index[i]
}
