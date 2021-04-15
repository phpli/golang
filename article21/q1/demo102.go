package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func String2ByteSlice(str string) (bs []byte) {
	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	sliceHdr.Data = strHdr.Data
	sliceHdr.Len = strHdr.Len
	sliceHdr.Cap = strHdr.Len
	return
}

func main()  {
	str := strings.Join([]string{"Go", "land"},"")
	s := String2ByteSlice(str)
	fmt.Printf("%s\n", s)
	s[5] = 'g'
	fmt.Println(str)
}
