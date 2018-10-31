package main

import "fmt"
import "strconv"

func main() {
	str, _ := strconv.ParseFloat("1.000", 64)
	fmt.Println(str)
}
