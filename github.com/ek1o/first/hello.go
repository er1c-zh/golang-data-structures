package main

import (
	"fmt"
	"github.com/ek1o/collections/BinaryTree"
	"math/rand"
)

func main() {
	root := BinaryTree.InitBinaryTree(4)
	for i := 0; i < 10; i++ {
		root.Insert(rand.Int() % 10)
	}
	fmt.Println(root)
	root = root.Del(4)
	fmt.Println(root)
}

