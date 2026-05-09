package main

import (
	"github.com/vearne/mem-align"
)

type Car struct {
	flag bool
	age  int32
	F1   int8
	F2   int64
	F3   *int32
	InnerStruct struct{
		InnerByte byte
		//InnerStr string
	}
	F4   []byte
	Name string
	F5   error
}

type ManyFields struct {
	F01 byte
	F02 byte
	F03 byte
	F04 byte
	F05 byte
	F06 byte
	F07 byte
	F08 byte
	F09 byte
	F10 byte
	F11 byte
	F12 byte
	F13 byte
	F14 byte
	F15 byte
	F16 byte
	F17 byte
	F18 byte
	F19 byte
	F20 byte
	F21 byte
	F22 byte
	F23 byte
	F24 byte
	F25 byte
	F26 byte
	F27 byte
	F28 byte
	F29 byte
	F30 byte
	F31 byte
}

func main() {
	memalign.PrintStructAlignment(ManyFields{})
	memalign.PrintStructAlignment(Car{})
}
