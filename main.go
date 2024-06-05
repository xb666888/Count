package main

import (
	"fmt"
	"homework/util"
)

func main() {
	fmt.Println(util.GetWordCount("hello world worldworld", "world"))
	fmt.Println(util.Tokenize("hello world world.World"))
}
