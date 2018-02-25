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
	bf int		// balance factory 定义为左子树高度减去右子树高度
}

func InitAVL(val int) (*Node) {
	return &Node {
		val,
		nil,
		nil,
		0,
	}
}

func (t *Node) Insert(val int) (*Node) {
	if t == nil {
		return &Node{
			val,
			nil,
			nil,
			0,
		}
	}

	if val > t.Val {
		// 插入到右子树
		if t.R == nil {
			// 右子树为空时
			t.R = &Node {
				val,
				nil,
				nil,
				0,
			}
			// 右子树从有到无 本节点的平衡因子减一
			t.bf--
		} else {
			// 右子树不为空时
			// 保存右子树的平衡因子
			rBfBeforeInsert := t.R.bf
			t.R = t.R.Insert(val)
			// 当右子树的平衡因子从0变化为1/-1 平衡因子减一
			if rBfBeforeInsert == 0 && math.Abs(float64(t.R.bf - rBfBeforeInsert)) > 0 {
				t.bf--
			}
		}

		if t.bf == -2 {
			if t.R.bf == -1 {
				// 刷新平衡因子
				t.bf = 0
				t.R.bf = 0
				// 左旋
				result := t.R
				t.R = result.L
				result.L = t
				return result
			} else {
				// 刷新平衡因子
				// 分析见表格
				t.R.L.bf = 0
				if t.R.L.bf == 1 {
					t.bf = 0
					t.R.bf = -1
				} else {
					t.bf = 1
					t.R.bf = 0
				}
				// 先在右子树上右旋
				newR := t.R.L
				t.R.L = newR.R
				newR.R = t.R
				t.R = newR
				// 再左旋
				result := t.R
				t.R = result.L
				result.L = t
				return result
			}
		}
	} else if val < t.Val {
		// 插入到左子树
		if t.L == nil {
			// 左子树为空
			t.L = &Node{
				val,
				nil,
				nil,
				0,
			}
			// 当左子树从无到有时 平衡因子加一
			t.bf++
		} else {
			lBFBeforeInsert := t.L.bf
			// 左子树不为空 递归插入
			t.L = t.L.Insert(val)
			// 当左子树平衡因子从0变成1/-1 平衡因子加一
			if lBFBeforeInsert == 0 && math.Abs(float64(t.L.bf - lBFBeforeInsert)) > 0 {
				t.bf++
			}
		}

		if t.bf == 2 {
			if t.L.bf == 1 {
				// 刷新平衡因子
				t.bf = 0
				t.L.bf = 0
				// 右旋
				result := t.L
				t.L = result.R
				result.R = t
				return result
			} else {
				if t.L.R.bf == 1 {
					// 刷新平衡因子
					t.bf = -1
					t.L.bf = 0
					t.L.R.bf = 0
				} else {
					// 刷新平衡因子
					t.bf = 0
					t.L.bf = 1
					t.L.R.bf = 0

				}
				// 先在左子树左旋
				newL := t.L.R
				t.L.R = newL.L
				newL.L = t.L
				t.L = newL
				// 再右旋
				result := t.L
				t.L = result.R
				result.R = t
				return result
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

// 中序遍历
func (t *Node) InOrder() string {
	if t != nil {
		var buf bytes.Buffer
		buf.WriteString(t.L.InOrder())
		buf.WriteString(strconv.Itoa(t.Val))
		buf.WriteString(" ")
		buf.WriteString(t.R.InOrder())
		return buf.String()
	}
	return ""
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
