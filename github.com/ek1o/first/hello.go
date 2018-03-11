package main

import (
	"github.com/ek1o/algorithm/string"
	"fmt"
)

func main() {
	fmt.Println(string.KMP("BBC ABCDAB ABCDABCDABD", "ABCDABD"))
}

