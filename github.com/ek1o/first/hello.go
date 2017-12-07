package main

import (
	"fmt"
	"github.com/ek1o/collections/BinaryTree"
)

func main() {
	root := BinaryTree.InitBinaryTree(4)
	root.Insert(1)
	root.Insert(6)
	root.Insert(-1)
	root.Insert(-2)
	fmt.Println(root)
}

