package main

const blockSize = 20
const dot = '.'
const hashTag = '#'
const newLine = '\n'

func isValidBlock(data []byte) bool {
	numDots, numHashtags := 0, 0
	for i := range data {
		switch data[i] {
		case dot:
			numDots++
		case hashTag:
			numHashtags++
		case newLine:
			if i%5 != 4 {
				return false
			}
		default:
			return false
		}
	}
	return numDots == 12 && numHashtags == 4
}

func validateFile(data []byte) (ok bool, res []tetrimino) {
	for i := 0; i < len(data); {
		if i+blockSize > len(data) {
			break
		}
		if !isValidBlock(data[i : i+blockSize]) {
			break
		}
		t := blockToTetrimino(data[i : i+blockSize])
		if !t.isValid() {
			break
		}
		res = append(res, t)
		i += blockSize
		if i == len(data) {
			ok = true
			break
		}
		if data[i] != '\n' {
			break
		}
		i += 1
	}
	return ok, res
}
