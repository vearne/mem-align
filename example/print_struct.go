package main

import (
	"github.com/vearne/mem-align"
)

type Car struct {
	flag bool
	age  int32
	F1   int8
	F2   int64
	InnerStruct struct{
		InnerByte byte
		//InnerStr string
	}
	F3   []byte
	Name string
	F4   error
}

func main(){
	memalign.PrintStructAlignment(Car{})
}
