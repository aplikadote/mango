package main

import "fmt"

import "github.com/aplikadote/mango/point"

func main() {
	fmt.Println("oli")
	v := point.New(1, 2)
	//v := point.Create()
	fmt.Println(v)
	fmt.Println(v.GetX())
	fmt.Println(v.GetY())
	fmt.Println(v.Z)
	fmt.Println(v.Abs())
}
