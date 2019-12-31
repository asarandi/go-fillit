package main

type tetrimino [4][]byte

var validUint16Map = map[uint16]bool{
	0x8888: true,
	0xf000: true,
	0xcc00: true,
	0xe400: true,
	0x4c40: true,
	0x4e00: true,
	0x8c80: true,
	0x88c0: true,
	0xe800: true,
	0xc440: true,
	0x2e00: true,
	0x44c0: true,
	0x8e00: true,
	0xc880: true,
	0xe200: true,
	0x6c00: true,
	0x8c40: true,
	0xc600: true,
	0x4c80: true,
}

func (t tetrimino) isEmptyRow(row int) bool {
	res := true
	for i := 0; i < 4; i++ {
		res = res && t[row][i] == dot
	}
	return res
}

func (t tetrimino) isEmptyColumn(column int) bool {
	res := true
	for i := 0; i < 4; i++ {
		res = res && t[i][column] == dot
	}
	return res
}

func (t tetrimino) toUint16() uint16 {
	var res uint16 = 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			res <<= 1
			if t[i][j] == hashTag {
				res |= 1
			}
		}
	}
	return res
}

func (t tetrimino) isValid() bool {
	return validUint16Map[t.toUint16()]
}

func blockToTetrimino(b []byte) tetrimino {
	in := tetrimino{b[0:4], b[5:9], b[10:14], b[15:19]}
	out := tetrimino{
		{dot, dot, dot, dot},
		{dot, dot, dot, dot},
		{dot, dot, dot, dot},
		{dot, dot, dot, dot},
	}
	y, x := 0, 0
	for in.isEmptyRow(y) {
		y++
	}
	for in.isEmptyColumn(x) {
		x++
	}
	for i := 0; i+y < 4; i++ {
		for j := 0; j+x < 4; j++ {
			out[i][j] = in[i+y][j+x]
		}
	}
	return out
}
