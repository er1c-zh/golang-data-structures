package AVL

import (
	"bytes"
	"strconv"
	"math"
)

type Node struct {
	Val int
	L *Node
	R *Node
}

func InitAVL(val int) (*Node) {
	return &Node {
		val,
		nil,
		nil,
	}
}

func (t *Node) Insert(val int) (*Node) {
	if t == nil {
		return &Node{
			val,
			nil,
			nil,
		}
	}

	if val > t.Val {
		t.R= t.R.Insert(val)
		hR := t.R.Height()
		hL := t.L.Height()
		if hR - hL == 2 {
			if val > t.R.Val {
				tmp := t.R.L
				result := t.R
				t.R.L = t
				t.R = tmp
				return result
			} else {
				rl := t.R.L
				t.R.L.R = t.R
				t.R.L = nil
				t.R = rl
				// right right 1
				tmp := t.R.L
				result := t.R
				t.R.L = t
				t.R = tmp
				return result
				// left 1
			}
		}
	} else if val < t.Val {
		t.L= t.L.Insert(val)
		hR := t.R.Height()
		hL := t.L.Height()
		if hL - hR == 2 {
			if val < t.L.Val {
				result := t.L
				tmp := t.L.R
				t.L.R = t
				t.L = tmp
				return result
			} else {
				lr := t.L.R
				t.L.R = nil
				lr.L = t.L
				t.L = lr
				// left left 1
				result := t.L
				tmp := t.L.R
				t.L.R = t
				t.L = tmp
				return result
				// right 1
			}
		}
	}
	return t
}

func (t *Node) Height() float64 {
	if t == nil {
		return 0
	}
	return 1 + math.Max(float64(t.L.Height()), float64(t.R.Height()))
}

func (t *Node) String() string {
	height := t.Height()
	var currentNodes, nextNodes []*Node
	var result []string
	currentNodes = append(currentNodes, t)
	isAllEmpty := false
	for !isAllEmpty {
		isAllEmpty = true
		var buf bytes.Buffer
		for i := range currentNodes {
			tmp := currentNodes[i]
			if tmp != nil {
				buf.WriteString(strconv.Itoa(currentNodes[i].Val))
				for j := 0; float64(j) < math.Pow(float64(2), float64(height)) / float64(len(currentNodes)); j++ {
					buf.WriteString("    ")
				}
				nextNodes = append(nextNodes, tmp.L)
				nextNodes = append(nextNodes, tmp.R)
				isAllEmpty = false
			} else {
				buf.WriteString(" ")
				for j := 0; float64(j) < math.Pow(float64(2), float64(height)) / float64(len(currentNodes)); j++ {
					buf.WriteString("    ")
				}
				nextNodes = append(nextNodes, nil)
				nextNodes = append(nextNodes, nil)
			}
		}
		result = append(result, buf.String())
		currentNodes = nextNodes
		nextNodes = nil
	}
	var buf bytes.Buffer
	for i := range result {
		for j := 0; float64(j) < math.Pow(float64(2), float64(len(result) - i)); j++ {
			buf.WriteString(" ")
		}
		buf.WriteString(result[i])
		buf.WriteString("\n")
	}
	return buf.String()
}
