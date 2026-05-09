package memalign

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
	"reflect"
	"strconv"
)

type FieldInfo struct {
	StrID string
	FieldType reflect.Type
	FieldName string
	FieldSize int
}

type ColorFunc func(format string, a ...interface{}) string
type MemCell struct {
	Marker     byte
	FieldIndex int
}

var (
	colorList []ColorFunc
)

func init() {
	colorList = []ColorFunc{
		color.RedString,
		color.GreenString,
		color.YellowString,
		color.BlueString,
		color.MagentaString,
		color.CyanString,
	}

}



func PrintStructAlignment(q interface{}) {
	buff := make([]MemCell, 0)
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		typeInfo := reflect.TypeOf(q)
		fieldInfoList := make([]FieldInfo, typeInfo.NumField())


		for i := 0; i < typeInfo.NumField(); i++ {
			fieldInfoList[i].StrID = string(fieldMarker(i))
			fieldInfoList[i].FieldType = typeInfo.Field(i).Type
			fieldInfoList[i].FieldName = typeInfo.Field(i).Name
		}

		for i := 0; i < typeInfo.NumField(); i++ {
			field := typeInfo.Field(i)
			size := int(field.Type.Size())
			fieldInfoList[i].FieldSize = size

			// padding
			paddSize := int(field.Offset) - len(buff)
			//fmt.Println(i, ":", field.Offset, len(buff)/2)
			for j := 0; j < paddSize; j++ {
				buff = append(buff, MemCell{Marker: ' ', FieldIndex: -1})
			}
			for k := 0; k < size; k++ {
				buff = append(buff, MemCell{
					Marker:     fieldMarker(i),
					FieldIndex: i,
				})
			}
		}

		// print
		// 1. Fields in struct
		PrintStructInfo(fieldInfoList)
		// 2. Memory layout
		ColorFormatPrint(buff)
	}

}

func fieldMarker(idx int) byte {
	return byte('A' + idx%26)
}

func PrintStructInfo(itemList []FieldInfo) {
	fmt.Println("---- Fields in struct ----")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.SetHeader([]string{"ID", "FieldType", "FieldName", "FieldSize"})

	for _, item := range itemList {
		table.Append([]string{item.StrID, item.FieldType.String(),
			item.FieldName, strconv.Itoa(item.FieldSize)})
	}
	table.Render() // Send output
}

func ColorFormatPrint(buff []MemCell) {
	fmt.Println("---- Memory layout ----")
	for i := 0; i < len(buff); i++ {
		fmt.Printf("|")
		if buff[i].FieldIndex >= 0 {
			idx := buff[i].FieldIndex % len(colorList)
			fmt.Printf(colorList[idx](string(buff[i].Marker)))
		} else {
			fmt.Printf(string(buff[i].Marker))
		}
		if (i+1)%8 == 0 {
			fmt.Println("|")
		}
	}
	if len(buff)%8 != 0 {
		remaining := 8 - len(buff)%8
		for i := 0; i < remaining; i++ {
			fmt.Printf("| ")
		}
		fmt.Println("|")
	}
	fmt.Printf("total cost: %v Bytes.\n", len(buff))
}