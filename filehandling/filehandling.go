package main

import (
	// "flag"
	"fmt"
	"io/ioutil"
	// "github.com/gobuffalo/packr"
)

func main() {
	//	box := packr.NewBox("../filehandling")
	//	data, _ := box.FindString("test.txt")
	//	fmt.Println("Contents of file:", data)

	//	fptr1 := flag.String("fpath", "test.txt", "Ubicacion del archivo")
	//	fptr2 := flag.String("fsep", ";", "Separador de campos")
	//	flag.Parse()
	//	fmt.Println("value of fpath is", *fptr1)
	//	fmt.Println("value of fsep is", *fptr2)

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
