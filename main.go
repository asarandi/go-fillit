package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const blocksize = 20

func tetriminoToUint16(b []byte) uint16 {
	var y,x,result = 0,0,0
	array := strings.Split(strings.TrimSpace(string(b)), "\n")
	//fixme

}

func isValidTetrimino(data []byte) bool {
	dots, hashtags := 0,0
	for i := range data {
		switch data[i] {
			case '.': dots += 1
			case '#': hashtags += 1
			case '\n': if i%5!=4 {return false}
			default: return false
		}
	}
	return dots==12 && hashtags==4
}

func isValid(data []byte) bool {
	i, result := 0, false
	for i<len(data) {
		if i+blocksize <= len(data) {
			if !isValidTetrimino(data[i:i+blocksize]) { break }
			if 0x29a == tetriminoToUint16(data[i:i+blocksize]) { break }
			i += blocksize
			if i == len(data) { result = true; break }
			if data[i] != '\n' { break }
			i += 1
		} else { break }
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./go-fillit filename")
		return
	}
	data, err := ioutil.ReadFile(os.Args[1]);
	if err != nil {
		fmt.Println("failed to read file")
		return
	}
	fmt.Println(isValid(data))
}
