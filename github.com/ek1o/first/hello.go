package main

import (
	"fmt"
	"github.com/ek1o/collections/AVL"
	"math/rand"
)

func main() {
	root := AVL.InitAVL(9)
	for i := 0; i < 15; i++ {
		fmt.Println("===")
		tmp := rand.Intn(100)
		fmt.Println(tmp)
		root = root.Insert(tmp)
		fmt.Println(root)
	}
	fmt.Println(root.InOrder())
}

