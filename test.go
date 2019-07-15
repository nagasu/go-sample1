package main

import "fmt"

// sample
type T struct {
	name  string // オブジェクトの名前
	value int    // その値
}

func test() {
	s := T{"test", 1}
	fmt.Println(s)
}
