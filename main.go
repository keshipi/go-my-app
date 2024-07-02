package main

import (
	"fmt"
)

func main() {
	// basic_()
	// map_()
	// const_()
	iota_()
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

func const_() {
	const n = 1
	x := 1 + n
	y := 1.2 + n
	fmt.Println(x, y)
}

func iota_() {
	const (
		Apple = iota
		Orange
		Banana
	)
	fmt.Println(Apple)

	const (
		Apple2 = iota + iota
		Orange2
		Banana2
	)
	fmt.Println(Apple2, Orange2, Banana2) // 0 2 4

	const (
		Apple3 = iota + iota
		Orange3
		Banana3 = iota + 3
	)
	fmt.Println(Apple3, Orange3, Banana3) // 0 2 5

	type Fruit int
	type Animal int
	const (
		Apple4 Fruit = iota
		Orange4
		Banana4
	)
	const (
		Monkey Animal = iota
		Elephant
		Pig
	)
	fmt.Println(Apple4, Orange4, Banana4) // 0 1 2
	fmt.Println(Monkey, Elephant, Pig)    // 0 1 2
}
