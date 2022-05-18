package main

import (
	"fmt"
	"io"
	"os"
)

// type content struct{}

// func (content) Write(bs []byte) (int, error) {
// 	fmt.Println("File contents:\n", string(bs))
// 	return len(bs), nil
// }

func main() {

	//fmt.Println(os.Args[1])
	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//fileContent := content{}

	//initially had implemented the commented out custom writer function and struct
	io.Copy(os.Stdout, file)
}
