package main

import (
	"fmt"

	a "github.com/akselsaatci/huffman/pkg/image_to_ascii"
)

func main() {
	string, err := a.ImageToAscii("/Users/akselsaatci/Developer/huffman_encoding/pkg/image_to_ascii/x.jpg")
	if err != nil {
		panic(err)
	}
	fmt.Print(string)
}
