# go-fillit

## usage

```sh
[pc:/tmp] asarandi% git clone -q https://github.com/asarandi/go-fillit
[pc:/tmp] asarandi% cd go-fillit/
[pc:/tmp/go-fillit] asarandi% go build
[pc:/tmp/go-fillit] asarandi% ./go-fillit samples/sample2.txt 
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
[pc:/tmp/go-fillit] asarandi% 
```


## tetriminos as `uint16`

### I
hex: **0x8888**

binary: `1000 1000 1000 1000 `

text:
```
#...
#...
#...
#...
```
---



### I
hex: **0xf000**

binary: `1111 0000 0000 0000 `

text:
```
####
....
....
....
```
---



### O
hex: **0xcc00**

binary: `1100 1100 0000 0000 `

text:
```
##..
##..
....
....
```
---



### T
hex: **0xe400**

binary: `1110 0100 0000 0000 `

text:
```
###.
.#..
....
....
```
---



### T
hex: **0x4c40**

binary: `0100 1100 0100 0000 `

text:
```
.#..
##..
.#..
....
```
---



### T
hex: **0x4e00**

binary: `0100 1110 0000 0000 `

text:
```
.#..
###.
....
....
```
---



### T
hex: **0x8c80**

binary: `1000 1100 1000 0000 `

text:
```
#...
##..
#...
....
```
---



### L
hex: **0x88c0**

binary: `1000 1000 1100 0000 `

text:
```
#...
#...
##..
....
```
---



### L
hex: **0xe800**

binary: `1110 1000 0000 0000 `

text:
```
###.
#...
....
....
```
---



### L
hex: **0xc440**

binary: `1100 0100 0100 0000 `

text:
```
##..
.#..
.#..
....
```
---



### L
hex: **0x2e00**

binary: `0010 1110 0000 0000 `

text:
```
..#.
###.
....
....
```
---



### J
hex: **0x44c0**

binary: `0100 0100 1100 0000 `

text:
```
.#..
.#..
##..
....
```
---



### J
hex: **0x8e00**

binary: `1000 1110 0000 0000 `

text:
```
#...
###.
....
....
```
---



### J
hex: **0xc880**

binary: `1100 1000 1000 0000 `

text:
```
##..
#...
#...
....
```
---



### J
hex: **0xe200**

binary: `1110 0010 0000 0000 `

text:
```
###.
..#.
....
....
```
---



### S
hex: **0x6c00**

binary: `0110 1100 0000 0000 `

text:
```
.##.
##..
....
....
```
---



### S
hex: **0x8c40**

binary: `1000 1100 0100 0000 `

text:
```
#...
##..
.#..
....
```
---



### Z
hex: **0xc600**

binary: `1100 0110 0000 0000 `

text:
```
##..
.##.
....
....
```
---



### Z
hex: **0x4c80**

binary: `0100 1100 1000 0000 `

text:
```
.#..
##..
#...
....
```
---




