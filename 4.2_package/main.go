package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type F struct {
	Name string
	Age  int
}

func (f *F) String() string {
	return fmt.Sprintf("NAME=%q, AGE=%d", f.Name, f.Age)
}

type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)

func (i Fruit) String() string {
	switch i {
	case Apple:
		return "Apple"
	case Orange:
		return "Orange"
	case Banana:
		return "Banana"
	}
	return fmt.Sprintf("Fruit(%d)", i)
}

var content = `
{
	"species": "はと",
	"description": "岩に灯るのが好き",
	"dimensions": {
		"height": 24,
		"width": 10
	}
}
`

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Data struct {
	Species     string     `json:"species"`
	Description string     `json:"description"`
	Dimensions  Dimensions `json:"dimensions"`
}

func main() {
	// f := &F{
	// 	Name: "John",
	// 	Age:  20,
	// }

	// fmt.Printf("%v\n", f)
	// fmt.Printf("%+v\n", f)
	// fmt.Printf("%#v\n", f)
	// fmt.Printf("%T\n", f)
	// fmt.Printf("%T\n", *f)

	var data Data
	err := json.Unmarshal([]byte(content), &data)
	if err != nil {
		log.Fatal(err)
	}
}
