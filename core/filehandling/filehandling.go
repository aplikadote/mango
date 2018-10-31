package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Test()
}

func Test() {
	filename := "D:\\makefile.txt"
	//	f, err := os.Create(filename)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("Oliwis")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
