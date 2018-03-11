package main

import (
	"fmt"
	"github.com/ek1o/algorithm/sort"
	"math/rand"
)

func main() {
	size := 1000
	foo := make([]int, size)
	for i := 0; i < size; i++ {
		foo[i] = rand.Intn(size)
	}
	fmt.Println("generate finish")
	fmt.Println(foo)
	sort.BubbleSort(foo, 0, size - 1)
	fmt.Println(foo)
	fmt.Println("sort finish")
	fmt.Println(size)
}

