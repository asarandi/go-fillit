package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const blockSize = 20
const dot = '.'
const hashTag = '#'
const newLine = '\n'

type tetrimino [4][]byte

var validUint16Map = map[uint16]bool{
	0x8888:true, 0xf000:true, 				                    // I
    0xcc00:true,         						                // O
    0xe400:true, 0x4c40:true, 0x4e00:true, 0x8c80:true, 		// T
	0x88c0:true, 0xe800:true, 0xc440:true, 0x2e00:true, 		// L
	0x44c0:true, 0x8e00:true, 0xc880:true, 0xe200:true,        	// J
	0x6c00:true, 0x8c40:true,                        	        // S
	0xc600:true, 0x4c80:true, 						            // Z
}

func (t tetrimino) isEmptyRow(row int) bool {
	result := true
	for i:=0; i<4; i++ { result = result && t[row][i] == dot }
	return result
}

func (t tetrimino) isEmptyColumn(column int) bool {
	result := true
	for i:=0; i<4; i++ { result = result && t[i][column] == dot }
	return result
}

func (t tetrimino) toUint16() uint16 {
	var result uint16 = 0
	for i:=0; i<4; i++ {
		for j:=0; j<4; j++ {
			result <<= 1
			if t[i][j] == hashTag { result |= 1 }
		}
	}
	return result
}

func (t tetrimino) isValid() bool {
	return validUint16Map[t.toUint16()]
}

func blockToTetrimino(b []byte) tetrimino {
	in := tetrimino{b[0:4], b[5:9], b[10:14], b[15:19]}
	out := tetrimino{
		{dot,dot,dot,dot},
		{dot,dot,dot,dot},
		{dot,dot,dot,dot},
		{dot,dot,dot,dot},
	}
	y,x := 0,0
	for in.isEmptyRow(y) { y++ }
	for in.isEmptyColumn(x) { x++ }
	for i:=0; i+y<4; i++ {
		for j:=0; j+x<4; j++ {
			out[i][j] = in[i+y][j+x]
		}
	}
	return out
}

func isValidBlock(data []byte) bool {
	numDots, numHashtags := 0,0
	for i := range data {
		switch data[i] {
			case dot: numDots++
			case hashTag: numHashtags++
			case newLine: if i%5!=4 {return false}
			default: return false
		}
	}
	return numDots==12 && numHashtags==4
}

func validateFile(data []byte) (ok bool, result []tetrimino) {
	for i:=0; i<len(data); {
		if i+blockSize > len(data) { break }
		if !isValidBlock(data[i:i+blockSize]) { break }
		t := blockToTetrimino(data[i:i+blockSize])
		if !t.isValid() { break }
		result = append(result, t)
		i += blockSize
		if i == len(data) { ok = true; break }
		if data[i] != '\n' { break }
		i += 1
	}
	return ok, result
}

type board [][]byte

func makeBoard(size int) board {
	var b board
	for i:=0;i<size;i++ {
		row := make([]byte, size)
		for j := range row {row[j] = dot}
		b = append(b, row)
	}
	return b
}

func (b board) print() {
	for i := range b {
		fmt.Println(string(b[i]))
	}
}

func (b board) check(i, j int, t tetrimino) bool {
	for y := range t {
		for x := range t[y] {
			if t[y][x] != hashTag { continue }
			if i+y >= len(b) || j+x >= len(b[i+y]) { return false }
			if b[i+y][j+x] != dot {	return false }
		}
	}
	return true
}

func (b board) put(i, j, idx int, t tetrimino) {
	for y := range t {
		for x := range t[y] {
			if t[y][x] != hashTag { continue }
			b[i+y][j+x] = byte(idx+'A')
		}
	}
}

func (b board) remove(i, j int, t tetrimino) {
	for y := range t {
		for x := range t[y] {
			if t[y][x] != hashTag { continue }
			b[i+y][j+x] = dot
		}
	}
}

func recursion(b board, array []tetrimino, idx int) bool {
	if idx == len(array) { return true }
	for i := range b {
		for j := range b[i] {
			if (b.check(i,j,array[idx])) {
				b.put(i, j, idx, array[idx]);
				if recursion(b, array, idx+1) { return true }
				b.remove(i, j, array[idx])
			}
		}
	}
	return false
}

func solve(array []tetrimino) board {
	square := 2
	for square*square < len(array) * 4 { square += 1}
	for {
		b := makeBoard(square)
		square += 1
		if !recursion(b, array, 0) { continue }
		return b
	}
}

func main() {
	if len(os.Args) != 2 { fmt.Println("usage: ./go-fillit filename"); return }
	data, err := ioutil.ReadFile(os.Args[1]);
	if err != nil {	fmt.Println("failed to read file"); return }
	ok, tetriminos := validateFile(data);
	if !ok { fmt.Println("error"); return }
	board := solve(tetriminos)
	board.print()
}
