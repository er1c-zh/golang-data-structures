package main

import (
	"fmt"
	"github.com/ek1o/collections/Heap"
	"math/rand"
)

func main() {
	root := Heap.InitHeap(15)
	for i := 0; i < 15; i++ {
		//fmt.Println("===")
		tmp := -rand.Intn(1000)
		fmt.Println(tmp)
		root.Insert(tmp)
		//isSuccess := root.Insert(tmp)
		//fmt.Println(isSuccess)
	}

	fmt.Println(root)

	for i := 0; i < 15; i++ {
		fmt.Println(root.PopMin())
	}
	//fmt.Println(root)
}

