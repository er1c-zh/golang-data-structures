package main

import (
	"fmt"
	"github.com/ek1o/golang-data-structures/algorithm/string"
)

func main() {
	fmt.Println(string.KMP("BBC ABCDAB ABCDABCDABD", "ABCDABD"))
}

