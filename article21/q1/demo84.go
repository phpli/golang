package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main()  {
	s := "颜色感染是一个有趣的游戏。"
	bs := []byte(s)
	s = string(bs)
	rs := []rune(s)
	s = string(rs)
	rs = bytes.Runes(bs)
	bs = Runes2Bytes(rs)
	fmt.Println(s,bs,s,rs,s,bs)
}

func Runes2Bytes(rs []rune)[]byte  {
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r)
	}
	n, bs := 0, make([]byte, n)
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
	}
	return bs
}