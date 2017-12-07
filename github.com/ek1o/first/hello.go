package main

import (
	"fmt"
	"github.com/ek1o/collections"
)

func main() {
	header := collections.InitLinkedList(100)
	header.Append(101).Append(102).Append(103).Append(101)
	fmt.Println(header.Find(101))
	header.DelAll(101)
	header.Insert(1, 0)
	fmt.Print(header.String())
}

