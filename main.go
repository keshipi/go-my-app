package main

import (
	"fmt"
	"sort"
)

func main() {
	// basic_()
	// map_()
	// const_()
	// iota_()
	// condition_()
	// array_()
	// string_()
	// type_()
	// struct_()
	v := Value(1)
	v = v.Add(2)
	fmt.Println(v)
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

// 順を持たないキーと値のペア
func map_() {
	// string型のキー、int型の値を持つmap定義
	var m map[string]int
	fmt.Println(m)
	// makeを使用して変数宣言も同時に行う
	m1 := make(map[string]int)
	m1["John"] = 21
	m1["Bob"] = 18
	m1["Mark"] = 33
	fmt.Println(m1)
	// リテラルを使って初期値を代入する
	m2 := map[string]int{
		"John": 21,
		"Bob":  33,
		"Mark": 33,
	}
	fmt.Println(m2)
	// 要素の削除
	delete(m2, "Bob")
	fmt.Println(m2)
	// mapのキーと値の列挙
	for k, v := range m2 {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
	// mapは順序を保持しないのでfor-rangeが毎回異なる結果になる。
	// ソートされてキーで列挙したい場合は先にキーを取り出し、ソート後にfor-rangeを実行する。
	keys := []string{}
	for k := range m1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)
	for _, k := range keys {
		fmt.Printf("key: %v, value: %v\n", k, m1[k])
	}
	// 存在しないキーは空文字になる
	m3 := map[string]string{
		"foo": "bar",
	}
	fmt.Println(m3["zoo"])
	// 存在したかのチェック
	v, ok := m3["foo"]
	if ok {
		fmt.Println(v)
	}
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

func condition_() {
	// if x == 2 && y == 3 {
	// 	fmt.Println("test")
	// }

	// if user, err := userName(); err == nil {
	// 	fmt.Println(user.Name)
	// } else {
	// 	log.Println(err)
	// }
}

func array_() {
	// int型 長さ4の配列: 固定長
	var a [4]int
	a[0] = 1
	fmt.Println(a)

	// スライスは長さを指定しない: 可変長
	var b []int
	// b[0] = 1 -> そのままアクセスするとpanicが発生する
	fmt.Println(b)

	// 変数宣言も兼ねた初期化
	c := make([]int, 3)
	c[0] = 1
	c[1] = 2
	c[2] = 3
	fmt.Println(c)
	// 初期値代入の初期化
	d := []int{1, 2, 3}
	fmt.Println(d)
	// スライスの伸長
	d = append(d, 4, 5, 6)
	fmt.Println(d, len(d), d[2:5])
	// スライスから要素の削除
	n := 1
	// d = append(d[:n], d[n+1:]...)
	// fmt.Println(d)
	d = d[:n+copy(a[n:], a[n+1:])]
	fmt.Println(d)
}

func string_() {
	// 文字列はバイト列で構成されているので、添え字を使ってアクセス
	s := "Hello"
	fmt.Printf("%c", s[0])
	// 内容を書き換えるにはバイト列に変換する
	b := []byte(s)
	b[0] = 'h'
	s = string(b)
	fmt.Printf("%c", s[0])
	// Unicodeのコードポイント列に変換する場合はrune型を使う
	s1 := "こんにちわ世界"
	rs := []rune(s1)
	rs[4] = 'は'
	s1 = string(rs)
	fmt.Println(s1)
	// ` で囲うことで複数行のテキストを扱う
	var content = `複数行の
	文章からなる
	テキストです。
	`
	fmt.Println(content)
}

func type_() {
	type MyString string
	var m MyString
	m = "foo"
	fmt.Println(m)
}

func struct_() {
	type User struct {
		Name string
		Age  int
	}
	var user User
	user.Name = "Bob"
	user.Age = 18
	// 変数宣言と初期値代入を同時に行う
	// user := User{
	// 	Name: "Bob",
	// 	Age:  18,
	// }
	fmt.Println(user)
}

type Value int

func (v Value) Add(n Value) Value {
	return v + n
}
