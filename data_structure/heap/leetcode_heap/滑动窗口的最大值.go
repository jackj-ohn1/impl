package leetcode_heap

import "impl/data_structure/heap"

func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	h := heap.NewHeap([]int{}, nums)

	for i := 0; i < len(nums); i++ {
		h.HeapPush(i)
		if i >= k-1 {
			for h.GetIndex(0) <= i-k {
				h.HeapPop()
			}
			ans = append(ans, h.GetValue(0))
		}
	}

	return ans
}
