package main

import (
	"fmt"
	"github.com/ek1o/collections/AVL"
)

func main() {
	root := AVL.InitAVL(0)
	for i := 1; i < 15; i++ {
		root = root.Insert(i)
	}
	fmt.Println(root)
}

