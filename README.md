Memory Alignment
=========
Memory-Alignment is a tool to help analyze layout of fields in struct in memory.

### Usage
```
go get github.com/vearne/mem-aligin
```

### Example
```
package main

import (
	"github.com/vearne/mem-align"
)

type Car struct {
	flag bool
	age  int32
	F1   int8
	F2   int64
	F3   []byte
	Name string
	F4   error
}

func main(){
	memalign.PrintStructAlignment(Car{})
}
```
### Output
![output](https://raw.githubusercontent.com/vearne/mem-align/master/output.jpeg)

