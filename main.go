package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./fillit filename")
		return
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("failed to read file")
		return
	}
	ok, tetriminos := validateFile(data)
	if !ok {
		fmt.Println("error")
		return
	}
	solution := solve(tetriminos)
	solution.print()
}
