package LinkedList

import (
	"bytes"
	"strconv"
)

/**
假定没有链表头
 */
type Node struct {
	val int
	next *Node
	pre *Node
}

func InitLinkedList(val int) *Node {
	return &Node {
		val,
		nil,
		nil,
	}
}

func (h *Node) String() string {
	var buf bytes.Buffer
	for h != nil {
		buf.WriteString(strconv.Itoa(h.val))
		if h.next != nil {
			buf.WriteString("->")
		}
		h = h.next
	}
	return buf.String()
}

func (h *Node) IsEmpty() bool {
	return h.next == nil
}

func (h *Node) Find(val int) (bool, *Node) {
	for h != nil {
		if h.val == val {
			return true, h
		}
		h = h.next
	}
	return false, nil
}

func (h *Node) IsLast() bool {
	return h.next == nil
}

func (h *Node) DelAll(val int) (count int, header *Node) {
	header = h
	count = 0
	if h == nil {
		return
	}
	for h.val == val {
		count++
		header = h.next
		h = h.next
	}
	for h != nil {
		if h.val == val {
			count++
			h.pre.next = h.next
			if h.next != nil {
				h.next.pre = h.pre
			}
			h = h.pre
		}
		h = h.next
	}
	return
}

func (h *Node) Insert(val int, pos int) (bool, *Node){
	cP := h
	len := 0
	if pos == 0 {
		new := &Node{h.val, h.next, h}
		h.val = val
		h.next = new
		if new.next != nil {
			new.pre = new
		}
		return true, h
	}
	for len < pos - 1 && cP != nil {
		cP = cP.next
		len++
	}
	if cP == nil {
		return false, h
	}
	new := &Node{val, cP.next, cP}
	if cP.next != nil {
		cP.next.pre = new
	}
	cP.next = new
	return true, h
}

func (h *Node) Append(val int) *Node {
	if h == nil {
		return InitLinkedList(val)
	}
	c := h
	for c.next != nil {
		c = c.next
	}
	c.next = &Node{
		val,
		nil,
		c,
	}
	return h
}
