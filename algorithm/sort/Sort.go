package sort

import (
	"github.com/ek1o/golang-data-structures/collections/Heap"
)

func BubbleSort(src []int, s int, e int) {
	for i := s ; i <= e ; i++ {
		for j := i; j <= e; j++ {
			if src[j] < src[i] {
				tmp := src[i]
				src[i] = src[j]
				src[j] = tmp
			}
		}
	}
}

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

func HeapSort(src []int, s int, e int) {
	heap := Heap.InitHeap(e - s + 1)
	for i := s; i <= e - s; i++ {
		heap.Insert(src[i])
	}
	for i := s; i <= e - s; i++ {
		src[i] = heap.PopMin()
	}
}

func MergeSort(src []int, s int, e int) {
	if e - s < 1 {
		return
	}
	if e - s == 1 {
		if src[e] < src[s] {
			tmp := src[e]
			src[e] = src[s]
			src[s] = tmp
		}
		return
	}

	mid := (s + e) / 2
	MergeSort(src, s, mid)
	MergeSort(src, mid + 1, e)
	tmp := make([]int, e - s + 1)

	i := s
	j := mid + 1
	for k := 0 ; (i <= mid || j <= e) && k <= e - s; k++ {
		if i <= mid && (j > e || src[i] < src[j]) {
			tmp[k] = src[i]
			i++
		} else {
			tmp[k] = src[j]
			j++
		}
	}
	for k := 0; k <= e - s; k++ {
		src[k + s] = tmp[k]
	}
}
