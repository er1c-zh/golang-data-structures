package Heap

import (
	"bytes"
	"strconv"
)

type Heap struct {
	pIdx int		// idx to insert
	heap []int
	length int
}

func InitHeap(length int) (*Heap) {
	 return &Heap {
		1,
		make([]int, length + 1),
		length,
	}
}

func (t *Heap) Insert(val int) bool {
	// 检查是否已满
	if t.pIdx > t.length {
		return false
	}
	// 上滤 直到上滤到root或者父节点比目标值小
	idx := t.pIdx
	for ; idx > 1 && t.heap[idx / 2] > val; idx = idx / 2 {
		t.heap[idx] = t.heap[idx / 2]
	}
	t.heap[idx] = val
	// 增加指针
	t.pIdx++
	return true
}

func (t *Heap) PopMin() int {
	// 空堆
	if t.pIdx < 2 {
		return -1
	}

	// 返回值 最小的值
	result := t.heap[1]

	// 维持堆序
	toInsert := t.heap[t.pIdx - 1]			// 最后一个值
	idx := 1
	// 下滤
	for ; idx * 2 < t.pIdx ; {
		// 查找最小子节点
		min := t.heap[idx * 2]
		minIdx := idx * 2
		if idx * 2 + 1 < t.pIdx && min > t.heap[idx * 2 + 1] {
			min = t.heap[idx * 2 + 1]
			minIdx = idx * 2 + 1
		}

		// 如果最小的细节点小于最后一个值 则继续
		if min < toInsert {
			t.heap[idx] = min
			idx = minIdx
		} else {
			// 大于的话就结束
			break
		}
	}
	// 插入最后一个值
	t.heap[idx] = toInsert
	t.pIdx--
	return result
}

func (t *Heap) String() string {
	var buf bytes.Buffer
	for i := range t.heap {
		buf.WriteString(strconv.Itoa(t.heap[i]))
		buf.WriteString(" ")
	}
	return buf.String()
}

