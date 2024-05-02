package leetcode_heap

import "impl/data_structure/heap"

func topKFrequent(nums []int, k int) []int {
	ref := map[int]int{}
	index := []int{}
	for i := 0; i < len(nums); i++ {
		ref[nums[i]]++
	}
	for k := range ref {
		index = append(index, k)
	}
	h := heap.NewHeap(index, ref)
	h.HeapInit()
	h.HeapSort()

	ans := []int{}
	n := h.Len() - 1
	for k > 0 {
		ans = append(ans, h.GetIndex(n))
		k--
		n--
	}
	return ans
}
