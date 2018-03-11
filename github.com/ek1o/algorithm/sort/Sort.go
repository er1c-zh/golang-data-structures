package sort

import (
	"github.com/ek1o/collections/Heap"
)

func QuickSort(src []int, s int, e int) {
	if e - s < 1 {
		return
	}

	i := s
	j := e - 1
	wait2swap := false
	target := src[e]

	for ; i < j; {
		if !wait2swap {
			if src[i] <= target {
				i++
			} else {
				wait2swap = true
			}
		} else {
			if src[j] > target {
				j--
			} else {
				tmp := src[j]
				src[j] = src[i]
				src[i] = tmp
				wait2swap = false
			}
		}
	}
	if target < src[i] {
		src[e] = src[i]
		src[i] = target
	} else {
		src[e] = src[i + 1]
		src[i + 1] = target
		i++
	}
	QuickSort(src, s, i - 1)
	QuickSort(src, i + 1, e)
}

func Heapsort(src []int, s int, e int) {
	heap := Heap.InitHeap(e - s + 1)
	for i := s; i <= e - s; i++ {
		heap.Insert(src[i])
	}
	for i := s; i <= e - s; i++ {
		src[i] = heap.PopMin()
	}
}

func Mergesort(src []int, s int, e int) {
}