package main

import "fmt"

func main() {
	basic_()
	map_()
}

func basic_() {
	// int型の変数宣言
	var n int
	n = 1
	fmt.Println(n)

	// varを省略した変数宣言
	x := 1
	fmt.Println(x)
}

func map_() {
	// string型のキー、int型の値を持つmap定義
	var m map[string]int
	fmt.Println(m)

	m1 := make(map[string]int)
	m1["John"] = 21
	m1["Bob"] = 18
	m1["Mark"] = 33
	fmt.Println(m1)

}
