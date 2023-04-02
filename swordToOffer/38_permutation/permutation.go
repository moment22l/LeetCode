package main

import "fmt"

var res []string
var byteMaps []map[byte]bool

func permutation(s string) []string {
	if len(s) == 0 {
		return nil
	}
	res = []string{}
	byteMaps = []map[byte]bool{}
	bs := []byte(s)
	for range bs {
		byteMaps = append(byteMaps, make(map[byte]bool))
	}
	for _, b := range bs {
		for i := 0; i < len(byteMaps); i++ {
			byteMaps[i][b] = false
		}
	}
	traceBack(bs, []byte{}, len(bs))
	return res
}

func traceBack(bs []byte, now []byte, length int) {
	if len(now) == length {
		temp := make([]byte, length)
		copy(temp, now)
		res = append(res, string(temp))
		return
	} else {
		for i := 0; i < len(bs); i++ {
			backup := make([]byte, len(bs))
			copy(backup, bs)
			b := bs[i]
			if byteMaps[len(now)][b] == false {
				bs = append(bs[:i], bs[i+1:]...)
				traceBack(bs, append(now, b), length)
				byteMaps[len(now)][b] = true
			}
			bs = backup
		}
		// 退出之前还原当前层的map
		for _, b := range bs {
			byteMaps[len(now)][b] = false
		}
	}
}

func main() {
	fmt.Println(permutation("a"))
}
